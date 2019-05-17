package main

import (
	"core/routes"
	"log"
	"os"
	"github.com/gin-contrib/cors"
)

func main() {


	f, err := os.OpenFile("core_log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("Server has started")

	//os.Chdir("/root/go/src/core")
	r := routes.SetupRouter()
	r.Use(cors.Default())
	//database.Connect, err = sql.Open(database.Driver, database.Dsn)
	if err != nil {
		log.Println(err)
	}
	r.Static("/static", "./files")
	r.Run(":4452")
	// Listen and Server in http://localhost:9119
	
}
