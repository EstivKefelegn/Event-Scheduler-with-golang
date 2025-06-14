package main

import (
	
	"Eventplanning.go/Api/db"
	"Eventplanning.go/Api/routes"
	"github.com/gin-gonic/gin"
)


func main()  {
	db.InitDB()
	
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}

