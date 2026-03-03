package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/packt-go-course/mytodo/models"
)

func ListTasks(c *gin.Context) {
	c.JSON(http.StatusOK, models.Tasks)

}
