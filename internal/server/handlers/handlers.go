package handlers

import (
	"github.com/amirmohammadkariimi/interview-task/pkg/cache"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
	c  *cache.Cache
}

func New(db *gorm.DB, c *cache.Cache) *Handler {
	return &Handler{
		db: db,
		c:  c,
	}
}
