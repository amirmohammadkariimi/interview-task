package middlewares

import (
	"strconv"
	"time"

	"github.com/amirmohammadkariimi/interview-task/internal/pkg/prometheus_metrics"
	"github.com/gin-gonic/gin"
)

func Prometheus() func(c *gin.Context) {
	return func(c *gin.Context) {
		start := time.Now()

		// Process the request
		c.Next()

		// Calculate request duration
		duration := time.Since(start).Seconds()

		// Collect HTTP status, method, and path
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		clientip := c.ClientIP()

		// Record metrics
		if status != 404 {
			prometheus_metrics.RequestDuration.WithLabelValues(path, method, strconv.Itoa(status)).Observe(duration)
		}
		prometheus_metrics.RequestCounter.WithLabelValues(path, method, strconv.Itoa(status), clientip).Inc()
	}
}
