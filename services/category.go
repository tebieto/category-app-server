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
	q := database.Conn.Create(&category)

	if q.Error != nil {
		log.Println("Error connecting to sql to store category with title" + category_title + "c.context is " )
		c.JSON(http.StatusBadRequest, gin.H{	
			"sql": q.Error,
		})
	
		return

	}

	log.Println("Category with title " + category_title + "was stored by c.context: " )


	c.JSON(http.StatusOK, gin.H{
		"id": category.Id,
		"title": category.Title,
		"slug":category.Slug,
		"parent": category.Parent,
	})

	return




}


// Read Category
func ReadCategory(c *gin.Context) {

	
	var Category []model.Category

	// Read
    q := database.Conn.Where("parent = ?", "0").Find(&Category)

	if q.Error != nil {
		log.Println("Error connecting to sql to read all categories" )
		c.JSON(http.StatusBadRequest, gin.H{	
			"sql": q.Error,
		})
	
		return

	}


	log.Println("Success reading all parent categories " )


	c.JSON(http.StatusOK, Category)

	return


}

// Read Category By Parent ID
func ReadCategoryByParentId(c *gin.Context) {


}

// Read Category By ID
func ReadCategoryById(c *gin.Context) {


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