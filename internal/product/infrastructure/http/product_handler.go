package http

import (
	"net/http"
	"strconv"
	"time"

	"prueba-tecnica-back/internal/product/application"
	"prueba-tecnica-back/internal/product/infrastructure/http/dto"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	createUseCase     *application.CreateProductUseCase
	listUseCase       *application.ListProductsUseCase
	updateUseCase     *application.UpdateProductUseCase
	deleteUseCase     *application.DeleteProductUseCase
	generateQRUseCase *application.GenerateQRUseCase
}

func NewProductHandler(
	createUseCase *application.CreateProductUseCase,
	listUseCase *application.ListProductsUseCase,
	updateUseCase *application.UpdateProductUseCase,
	deleteUseCase *application.DeleteProductUseCase,
	generateQRUseCase *application.GenerateQRUseCase,
) *ProductHandler {
	return &ProductHandler{
		createUseCase:     createUseCase,
		listUseCase:       listUseCase,
		updateUseCase:     updateUseCase,
		deleteUseCase:     deleteUseCase,
		generateQRUseCase: generateQRUseCase,
	}
}

func (h *ProductHandler) Create(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	product, err := h.createUseCase.Execute(req.Name, req.Category, req.Price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := dto.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Category:  product.Category,
		Price:     product.Price,
		QRCode:    product.QRCode,
		CreatedAt: product.CreatedAt.Format(time.RFC3339),
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *ProductHandler) List(c *gin.Context) {
	var query dto.ListProductsQuery
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

	products, total, err := h.listUseCase.Execute(query.Category, query.SortOrder, query.Page, query.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch products"})
		return
	}

	var resp []dto.ProductResponse
	for _, prod := range products {
		resp = append(resp, dto.ProductResponse{
			ID:        prod.ID,
			Name:      prod.Name,
			Category:  prod.Category,
			Price:     prod.Price,
			QRCode:    prod.QRCode,
			CreatedAt: prod.CreatedAt.Format(time.RFC3339),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"products": resp,
		"total":    total,
		"page":     query.Page,
		"limit":    query.Limit,
	})
}

func (h *ProductHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	product, err := h.updateUseCase.Execute(uint(id), req.Name, req.Category, req.Price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := dto.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Category:  product.Category,
		Price:     product.Price,
		QRCode:    product.QRCode,
		CreatedAt: product.CreatedAt.Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	err = h.deleteUseCase.Execute(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
