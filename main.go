package main

import (
	"github.com/gin-gonic/gin"
	"weBook/internal/web/user"
)

func main() {
	server := gin.Default()
	// server.Use(utils.Cors())
	u := &user.UserHandler{}
	u.RegisterRoutesUser(server)
	server.Run("localhost:8080")
}
