package handlers

import (
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/amirmohammadkariimi/interview-task/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Lookup() func(c *gin.Context) {
	return func(c *gin.Context) {
		domain := c.Query("domain")
		if domain == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "'domain' query parameter is required",
			})
			return
		}

		q, found := h.c.Get(domain)
		if found {
			if q.CreatedAt > time.Now().Add(-1*time.Hour).Unix() {
				c.Header("x-cache", "HIT")
				c.JSON(http.StatusOK, q)
				return
			} else {
				h.c.Delete(domain)
			}
		}
		ips, err := net.LookupIP(domain)
		if err != nil {
			if dnsErr, ok := err.(*net.DNSError); ok && dnsErr.IsNotFound {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "domain not found",
				})
				return
			} else {
				slog.Error("Could not get IPs for domain:", domain, err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "server error, please contact website administrator",
				})
				return
			}
		}
		var addresses []models.Address
		for _, ip := range ips {
			addresses = append(addresses, models.Address{
				IP: ip.String(),
			})
		}
		query := models.DNSQuery{
			ClientIp:  c.ClientIP(),
			Domain:    domain,
			CreatedAt: time.Now().Unix(),
			Addresses: addresses,
		}
		h.c.Set(domain, query)
		result := h.db.Create(&query)
		if result.Error != nil {
			slog.Error("error saving query to database", "error", result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "server error, please contact website administrator",
			})
			return
		}
		c.Header("x-cache", "MISS")
		c.JSON(http.StatusOK, query)
	}
}
