package database

import (
	"core/models"
	//"fmt"
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//Conn exported
var Conn *gorm.DB

//TODO panic when conn is nil
func init() {

	db := new(Config)

	available_db := Available()

	db.Host = available_db.Host
	db.User = available_db.User
	db.Password = available_db.Password
	db.Database = available_db.Database
	
	conn, err := gorm.Open("mysql", db.User+":"+db.Password+
		"@/"+db.Database+"?charset=utf8")
	if err != nil {
		panic(err)

	}

	
	 m := conn.AutoMigrate(&model.Category{})
	 
	 if m.Error != nil {

		panic(err)

	 }
	
	 if err != nil {
		panic(err)

	}

	Conn = conn
}