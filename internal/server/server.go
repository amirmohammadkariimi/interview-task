package server

import (
	"github.com/amirmohammadkariimi/interview-task/internal/server/handlers"
	"github.com/amirmohammadkariimi/interview-task/internal/server/middlewares"
	"github.com/amirmohammadkariimi/interview-task/pkg/cache"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/gorm"
)

type Server struct {
	e    *gin.Engine
	port string
	h    *handlers.Handler
}

func New(port string, db *gorm.DB, c *cache.Cache) *Server {
	return &Server{
		e:    gin.Default(),
		port: port,
		h:    handlers.New(db, c),
	}
}

func (s *Server) Run() error {
	s.registerMiddlewares()
	s.registerRoutes()
	return s.e.Run(s.port)
}

func (s *Server) registerRoutes() {
	s.e.GET("/metrics", gin.WrapH(promhttp.Handler()))
	s.e.GET("/health", s.h.Health())
	s.e.GET("/", s.h.Root())
	s.e.GET("/v1/tools/lookup", s.h.Lookup())
	s.e.GET("/v1/tools/validate", s.h.Validate())
	s.e.GET("/v1/history", s.h.History())
}

func (s *Server) registerMiddlewares() {
	s.e.Use(middlewares.Prometheus())
}
