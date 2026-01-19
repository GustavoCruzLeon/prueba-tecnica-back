package http

import (
	"net/http"
	"strconv"
	"time"

	"prueba-tecnica-back/internal/customer/application"
	"prueba-tecnica-back/internal/customer/infrastructure/http/dto"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	createUseCase *application.CreateCustomerUseCase
	listUseCase   *application.ListCustomersUseCase
	updateUseCase *application.UpdateCustomerUseCase
	deleteUseCase *application.DeleteCustomerUseCase
}

func NewCustomerHandler(
	createUseCase *application.CreateCustomerUseCase,
	listUseCase *application.ListCustomersUseCase,
	updateUseCase *application.UpdateCustomerUseCase,
	deleteUseCase *application.DeleteCustomerUseCase,
) *CustomerHandler {
	return &CustomerHandler{
		createUseCase: createUseCase,
		listUseCase:   listUseCase,
		updateUseCase: updateUseCase,
		deleteUseCase: deleteUseCase,
	}
}

func (h *CustomerHandler) Create(c *gin.Context) {
	var req dto.CreateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	customer, err := h.createUseCase.Execute(req.Name, req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := dto.CustomerResponse{
		ID:        customer.ID,
		Name:      customer.Name,
		Email:     customer.Email,
		CreatedAt: customer.CreatedAt.Format(time.RFC3339),
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *CustomerHandler) List(c *gin.Context) {
	var query dto.ListCustomersQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query parameters"})
		return
	}

	if query.Page < 1 {
		query.Page = 1
	}
	if query.Limit < 1 {
		query.Limit = 10
	}
	if query.Limit > 100 {
		query.Limit = 100
	}

	customers, total, err := h.listUseCase.Execute(query.Name, query.Page, query.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch customers"})
		return
	}

	var resp []dto.CustomerResponse
	for _, cust := range customers {
		resp = append(resp, dto.CustomerResponse{
			ID:        cust.ID,
			Name:      cust.Name,
			Email:     cust.Email,
			CreatedAt: cust.CreatedAt.Format(time.RFC3339),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"customers": resp,
		"total":     total,
		"page":      query.Page,
		"limit":     query.Limit,
	})
}

func (h *CustomerHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer ID"})
		return
	}

	var req dto.CreateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	customer, err := h.updateUseCase.Execute(uint(id), req.Name, req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := dto.CustomerResponse{
		ID:        customer.ID,
		Name:      customer.Name,
		Email:     customer.Email,
		CreatedAt: customer.CreatedAt.Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, resp)
}

func (h *CustomerHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer ID"})
		return
	}

	err = h.deleteUseCase.Execute(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
