package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	Handler struct {
		db *gorm.DB
	}
)

func NewHandler(db *gorm.DB) (h *Handler, err error) {
	h = &Handler{
		db: db,
	}
	return
}

func (h *Handler) rootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func (h *Handler) authRootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}
