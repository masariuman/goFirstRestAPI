package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/masariuman/goFirstRestAPI/controllers/productcontroller"
)

func productRoute(superRoute *gin.RouterGroup) {
	productsRoute := superRoute.Group("/product")
	{
		productsRoute.GET("/", productcontroller.Index)
		productsRoute.GET("/:id", productcontroller.Show)
		productsRoute.POST("", productcontroller.Store)
		productsRoute.PUT("/:id", productcontroller.Update)
		productsRoute.DELETE("/", productcontroller.Delete)
	}
}
