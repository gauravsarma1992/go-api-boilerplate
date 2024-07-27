package server

import (
	"log"

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

func LogError(err error) {
	if err != nil {
		log.Println("Error in base_handlers", err)
	}
}

func HandleAuthErrorJSON(c *gin.Context, err error) {
	HandleJSON(c, 401, gin.H{
		"message": "You are not allowed to access the API/resource",
	}, err)
}

func HandleClientErrorJSON(c *gin.Context, v interface{}, err error) {
	HandleJSON(c, 400, gin.H{
		"message": "failed",
		"reason":  v,
	}, err)
}

func HandleServerErrorJSON(c *gin.Context, v interface{}, err error) {
	HandleJSON(c, 500, gin.H{
		"message": "failed",
		"reason":  v,
	}, err)
}

func HandleSuccessJSON(c *gin.Context, v interface{}) {
	HandleJSON(c, 200, gin.H{
		"message": "success",
		"data":    v,
	}, nil)
}

func HandleJSON(c *gin.Context, status int, v interface{}, err error) {
	LogError(err)
	c.JSON(status, v)
}
