package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/packt-go-course/mytodo/handlers"
)

func UpdateRoutes(router *gin.Engine) {
	router.PUT("/tasks/:id", handlers.UpdateTask)
}
