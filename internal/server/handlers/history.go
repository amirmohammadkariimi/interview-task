package handlers

import (
	"net/http"

	"github.com/amirmohammadkariimi/interview-task/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) History() func(c *gin.Context) {
	return func(c *gin.Context) {
		var dnsQueries []models.DNSQuery
		h.db.Preload("Addresses").Order("created_at DESC").Limit(20).Find(&dnsQueries)
		c.JSON(http.StatusOK, dnsQueries)
	}
}
