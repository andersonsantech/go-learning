package main

import (
	"log"
	"github.com/andersonsantech/go-learning/internal/container"
	"github.com/gin-gonic/gin"
)

const Port = ":8080"

func main() {
	startServer()
}

func startServer() {
	container := container.InjectUserDependencies()

	server := gin.Default()

	server.Use(cors())

	server.POST("/users", func(c *gin.Context) {
		container.UserHandler.RegisterUser(c.Writer, c.Request)
	})

	server.GET("/users/:id", func(c *gin.Context) {
		container.UserHandler.GetUserByID(c.Writer, c.Request, c.Params.ByName("id"))
	})

	server.GET("/users", func(c *gin.Context) {
		container.UserHandler.FindAllUsers(c.Writer, c.Request)
	})

	server.DELETE("/users/:id", func(c *gin.Context) {
		container.UserHandler.DeleteUser(c.Writer, c.Request, c.Params.ByName("id"))
	})

	if err := server.Run(Port); err != nil {
		log.Fatal(err)
	}

}


func cors() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}