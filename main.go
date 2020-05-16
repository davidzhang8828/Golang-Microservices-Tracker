package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	// set up default Gin router
	router := gin.Default()

	// server up static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// set up route group
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	// start server
	router.Run(":3000")
}
