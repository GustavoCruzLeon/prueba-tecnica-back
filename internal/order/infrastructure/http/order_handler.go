package http

import (
	"net/http"
	"strconv"
	"time"

	"prueba-tecnica-back/internal/order/application"
	"prueba-tecnica-back/internal/order/domain"
	"prueba-tecnica-back/internal/order/infrastructure/http/dto"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	createUseCase         *application.CreateOrderUseCase
	listOrdersUseCase     *application.ListOrdersByCustomerUseCase
	calculateTotalUseCase *application.CalculateTotalSpentUseCase
}

func NewOrderHandler(
	createUseCase *application.CreateOrderUseCase,
	listOrdersUseCase *application.ListOrdersByCustomerUseCase,
	calculateTotalUseCase *application.CalculateTotalSpentUseCase,
) *OrderHandler {
	return &OrderHandler{
		createUseCase:         createUseCase,
		listOrdersUseCase:     listOrdersUseCase,
		calculateTotalUseCase: calculateTotalUseCase,
	}
}

func (h *OrderHandler) Create(c *gin.Context) {
	var req dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Convertir DTOs a entidades del dominio
	items := make([]domain.OrderItem, len(req.Items))
	for i, item := range req.Items {
		items[i] = domain.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
	}

	order, err := h.createUseCase.Execute(req.CustomerID, items)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convertir respuesta a DTO
	resp := convertToOrderResponse(order)
	c.JSON(http.StatusCreated, resp)
}

func (h *OrderHandler) ListByCustomer(c *gin.Context) {
	customerIDParam := c.Param("customer_id")
	customerID, err := strconv.ParseUint(customerIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer ID"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	orders, total, err := h.listOrdersUseCase.Execute(uint(customerID), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch orders"})
		return
	}

	// Convertir todas las órdenes a DTOs
	resp := make([]dto.OrderResponse, len(orders))
	for i, order := range orders {
		resp[i] = convertToOrderResponse(&order)
	}

	c.JSON(http.StatusOK, gin.H{
		"orders": resp,
		"total":  total,
		"page":   page,
		"limit":  limit,
	})
}

func (h *OrderHandler) CalculateTotalSpent(c *gin.Context) {
	customerIDParam := c.Param("customer_id")
	customerID, err := strconv.ParseUint(customerIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer ID"})
		return
	}

	total, err := h.calculateTotalUseCase.Execute(uint(customerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to calculate total"})
		return
	}

	c.JSON(http.StatusOK, dto.TotalSpentResponse{
		CustomerID: uint(customerID),
		Total:      total,
	})
}

func convertToOrderResponse(order *domain.Order) dto.OrderResponse {
	var shippedAtStr *string
	if order.ShippedAt != nil {
		str := order.ShippedAt.Format(time.RFC3339)
		shippedAtStr = &str
	}

	createdAtStr := order.CreatedAt.Format(time.RFC3339)

	var customer *dto.Customer
	if order.Customer != nil {
		customer = &dto.Customer{
			ID:    order.Customer.ID,
			Name:  order.Customer.Name,
			Email: order.Customer.Email,
		}
	}

	items := make([]dto.OrderItem, len(order.Items))
	for i, item := range order.Items {
		items[i] = dto.OrderItem{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
	}

	return dto.OrderResponse{
		ID:              order.ID,
		Status:          order.Status,
		Total:           order.Total,
		ShippingAddress: order.ShippingAddress,
		ShippedAt:       shippedAtStr,
		CustomerID:      order.CustomerID,
		Customer:        customer,
		Items:           items,
		CreatedAt:       createdAtStr,
	}
}
