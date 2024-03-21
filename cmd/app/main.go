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
