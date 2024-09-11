package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/amirmohammadkariimi/interview-task/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

var (
	version = "0.0.1"
)

func (h *Handler) Root() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, models.RootHandlerResponse{
			Date:       time.Now().Unix(),
			Kubernetes: isRunningInKubernetes(),
			Version:    version,
		})
	}
}

// Function to check if the app is running inside Kubernetes
func isRunningInKubernetes() bool {
	_, exists := os.LookupEnv("KUBERNETES_SERVICE_HOST")
	return exists
}
