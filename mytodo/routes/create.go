package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/packt-go-course/mytodo/handlers"
)

func CreateRoutes(router *gin.Engine) {
	router.POST("/tasks", handlers.CreateTask)
}
