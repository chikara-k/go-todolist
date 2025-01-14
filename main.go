package main

import (
	"log"
	"net/http"

	"github.com/chikara-k/go-todolist/db"
	"github.com/chikara-k/go-todolist/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db, err := db.ConnectionDB()
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	//r := api.NewRouter(db)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
