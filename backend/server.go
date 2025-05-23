package main

import (
	"fmt"
	"github.com/gin-gonic/gin" 
	"github.com/go_backend/routes"
)

func main() {
	fmt.Println("Starting the backend...")
	router := gin.Default()

	router.GET("/notes", routes.GetNote)
	router.Run(":8080")
}