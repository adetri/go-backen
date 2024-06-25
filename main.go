package main

import (
	"testbe/api"
	_ "testbe/docs" // This imports the generated Swagger docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	// Set up the Swagger handler
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Set up the API routes with the new base path
	r.GET("/hello", api.HelloHandler)

	r.Run(":8080")
}
