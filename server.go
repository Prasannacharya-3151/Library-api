package main

import (
	"library-api/config"
	"library-api/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No env file found in the system vars")
	}

	config.ConnectDB()
	config.MigrateDB() //this means no need create a table manually it will create tables automattically

	r := gin.Default() //this will create server. r is the http server inside router, recovery, router, middleware everything get initialized and then gin.Default() is the engine get ready eccept the request
	routes.SetupRouter(r)

	port := os.Getenv("PORT")
	if port == ""{
		port = "8018"
	}

	log.Printf("server is running on port %s", port)
	r.Run(":" + port)
}