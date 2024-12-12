package main

import (
	"HD-Wallet/db"
	"HD-Wallet/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	server := gin.Default()
	routes.RegisteredRoutes(server)
	server.Run(":8080")
}
