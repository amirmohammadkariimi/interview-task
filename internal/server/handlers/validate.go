package handlers

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Validate() func(c *gin.Context) {
	return func(c *gin.Context) {
		ip := c.Query("ip")
		if ip == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "'ip' query parameter is required",
			})
			return
		}
		parsedIp := net.ParseIP(ip)
		if parsedIp == nil {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": true,
		})
	}
}
