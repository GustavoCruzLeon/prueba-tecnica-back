package http

import (
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"

	"prueba-tecnica-back/internal/product/application"
	"prueba-tecnica-back/internal/product/infrastructure/http/dto"

	"github.com/gin-gonic/gin"
)

type QRHandler struct {
	generateQRUseCase *application.GenerateQRUseCase
}

func NewQRHandler(generateQRUseCase *application.GenerateQRUseCase) *QRHandler {
	return &QRHandler{
		generateQRUseCase: generateQRUseCase,
	}
}

func (h *QRHandler) GenerateQR(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	product, err := h.generateQRUseCase.Execute(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Category:  product.Category,
		Price:     product.Price,
		QRCode:    product.QRCode,
		CreatedAt: product.CreatedAt.Format("2006-01-02T15:04:05Z"),
	})
}

func (h *QRHandler) GetQRImage(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	product, err := h.generateQRUseCase.Execute(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qrData := product.QRCode
	if strings.HasPrefix(qrData, "data:image/png;base64,") {
		qrData = strings.SplitAfter(qrData, "data:image/png;base64,")[1]
	}

	imgBytes, err := base64.StdEncoding.DecodeString(qrData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode QR image"})
		return
	}

	c.Header("Content-Type", "image/png")
	c.Writer.Write(imgBytes)
}
