package routes

import (
	"net/http"
	"core/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//SetupRouter handles routes of the system
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	//Route for Default Server
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the main GPI EDUC8 Web Server. Powered by GO!!!")
	})

	// Category route
	r.POST("/category", func(c *gin.Context) {
		controllers.Category(c)
	})


	return r
}
