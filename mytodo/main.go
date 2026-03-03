package main

import (
	"github.com/gin-gonic/gin"
	"github.com/packt-go-course/mytodo/routes"
)

func main() {
	router := gin.Default()
	routes.CreateRoutes(router)
	routes.UpdateRoutes(router)
	routes.DeleteRoutes(router)
	routes.ListRoutes(router)

	router.Run(":8080")
}
