package main

import (
	"library-api/config"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := godotenv(); err != nil {
		log.Println("No env file found in the system vars")
	}

	config.ConnectDB()
	config.Migrate() //this measn no need create a table manually it will create tables automattically

	r := gin.Default() //this will create server. r is the http server inside router, recovery, router, middleware everything get initialized and then gin.Default() is the engine get ready eccept the request
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == ""{
		port = "8018"
	}

	log.Printf("server is running on port %s", port)
	r.Run(":" + port)
}