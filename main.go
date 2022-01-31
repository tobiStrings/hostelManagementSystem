package main

import (
	"github.com/gin-gonic/gin"
	"hostelManagementSystem/routes"
	"os"
)

func main() {
	os.Setenv("PORT","9090")
	port:= os.Getenv("PORT")

	if port == ""{
		port = "9090"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.StudentRoute(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200,gin.H{"success":"Access granted for api-1"})
	})

	router.Run( ":"+port)
}
