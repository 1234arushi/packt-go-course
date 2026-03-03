package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/packt-go-course/mytodo/handlers"
)

func DeleteRoutes(router *gin.Engine) {
	router.DELETE("/tasks/:id", handlers.DeleteTask)
}
