package main

import (
	"log"

	"github.com/gin-gonic/gin"

	controllers "ChGo/controllers"
	"ChGo/db"
	middlewares "ChGo/middlewares"
)

func main() {
	log.Println("Starting server..")

	db.Init()

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		product := v1.Group("/product", middlewares.Auth)
		{
			product.GET("/", controllers.GetProducts)
			product.POST("/", controllers.CreateProduct)
			product.PUT("/:id", controllers.UpdateProduct)
			product.DELETE("/:id", controllers.DeleteProduct)
		}

		noauth := v1.Group("/noauth")
		{
			noauth.POST("/login", middlewares.Sign)
			noauth.POST("/register", controllers.Register)
			noauth.GET("/refresh", middlewares.Refresh)
		}

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.Run()
}
