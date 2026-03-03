package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/packt-go-course/mytodo/handlers"
)

func ListRoutes(router *gin.Engine) {
	router.GET("/tasks", handlers.ListTasks)
}
