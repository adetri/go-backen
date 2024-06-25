package api

import (
	"github.com/gin-gonic/gin"
)

// HelloHandler godoc
// @Summary Show a message
// @Description get string message
// @Tags example
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /hello [get]
func HelloHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello, World!"})
}
