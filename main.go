package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = gin.Accounts{
	"ardimr": "ardi123",
	"admin":  "admin123",
}

func main() {

	server := gin.New()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	adminRouter := server.Group("/admin")
	adminRouter.Use(Authorize)

	adminRouter.GET("/test", TestFunc)

	server.POST("/login", Authenticate)

	server.GET("/redirect", Redirect)
	server.Run("localhost:8000")
}

func TestFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Message": "OK"})
}
func Redirect(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Message": "Redirected to here"})
}
