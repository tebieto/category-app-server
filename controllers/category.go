package controllers

import (
	"net/http"
	//"core/models"
	"core/services"
	//"time"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Category controller
func Category(c *gin.Context) {

	crud_action := c.PostForm("action")

	fmt.Println(crud_action)
	
	if (crud_action == "create") {
		services.CreateCategory(c)
	} else if (crud_action == "read") {
		services.ReadCategory(c)
	} else if (crud_action == "update") {
		services.UpdateCategory(c)
	} else if (crud_action == "delete") {
		services.DeleteCategory(c)
	} else {

		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
		})
		
		return

	}




}