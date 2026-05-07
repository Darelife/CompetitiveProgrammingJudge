package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"judge/internal/db"
	"judge/internal/transport"
	"judge/internal/worker"
)

func main() {
	err := db.Init()
	if err != nil {
		log.Fatal(err)
	}

	worker.StartWorker()

	router := gin.Default()

	transport.RegisterRoutes(router)

	log.Println("Server started on :4000")

	router.Run(":4000")
}
