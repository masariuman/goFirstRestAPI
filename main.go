package main

import (
	"github.com/gin-gonic/gin"
	"github.com/masariuman/goFirstRestAPI/models"
	"github.com/masariuman/goFirstRestAPI/routes"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	router := r.Group("/api")
	routes.AddRoutes(router)

	r.Run()
}
