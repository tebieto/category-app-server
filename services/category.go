package services

import (
	"log"
	"net/http"
	"core/models"
	"core/database"
	//"time"
	//"golang.org/x/crypto/bcrypt"
	//jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"strconv"
	//"strings"
	//"fmt"
	//"math/rand"
)

// Create Category
func CreateCategory(c *gin.Context) {

	category_title := c.PostForm("category_title")
	category_parent, _ := strconv.Atoi(c.PostForm("category_parent"))
	category_slug := slug.Make(category_title)

	// Create
	var category = model.Category{Title: category_title, Parent: category_parent, Slug: category_slug}
	m := database.Conn.Create(&category)

	if m.Error != nil {
		log.Println("Error connecting to sql to store category with title" + category_title + "c.context is " )
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": m.Error,
		})
	
		return

	}

	log.Println("Category with title " + category_title + "was stored by c.context: " )


	c.JSON(http.StatusOK, gin.H{
		"Success": category.Id,
	})

	return




}


// Read Category
func ReadCategory(c *gin.Context) {

	//category_id := c.PostForm("category_id")

}


// Update Category
func UpdateCategory(c *gin.Context) {

	//category_id := c.PostForm("category_id")
	//category_title := c.PostForm("category_title")

}


// Delete Category
func DeleteCategory(c *gin.Context) {

	//category_id := c.PostForm("category_id")

}