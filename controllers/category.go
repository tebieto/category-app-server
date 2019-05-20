package controllers

import (
	"net/http"
	//"core/models"
	"core/services"
	//"time"
	//"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
)

// Category controller
func Category(c *gin.Context) {

	crud_action := c.PostForm("action")
	category_parent, _ := strconv.Atoi(c.PostForm("category_parent"))
	category_id, _ := strconv.Atoi(c.PostForm("category_id"))

	if (crud_action == "create") {
		services.CreateCategory(c)

		return
	} 
	
	if (crud_action == "read") {
		
		// If you want to read details of a single category by their ID
		if (category_id>0){

		services.ReadCategoryById(c)

		return

		}
		
		// If you want details of more categories

		// If you did not specify a parent

	  if (category_parent == 0){

			services.ReadCategory(c)
			return
	
		} 
		
		// If you specify a parent
	  if (category_parent>0){

			services.ReadCategoryByParentId(c)
			return
		
			}
			

			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Bad Request",
			})
			
			return
	
	}

	
	 if (crud_action == "update") {
		services.UpdateCategory(c)
		return
	} 
	
	if (crud_action == "delete") {
		services.DeleteCategory(c)
		return
	} 
	


	c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
		})
		
		return





}