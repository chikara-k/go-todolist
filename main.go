package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/chikara-k/go-todolist/db"
	"github.com/chikara-k/go-todolist/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func listeners(r *gin.Engine, dbConnection *gorm.DB) {
	r.POST("/todo/create", func(c *gin.Context) {
		content := c.PostForm("content")
		fmt.Println(c.Request.PostForm, content)
		result := dbConnection.Create(&models.Todo{Content: content})
		if db.ErrorDB(result, c) {
			return
		}
	})

	r.GET("/todo/list", func(c *gin.Context) {
		var todos []models.Todo
		// Get all records
		result := dbConnection.Find(&todos)
		if db.ErrorDB(result, c) {
			return
		}
		fmt.Println(json.NewEncoder(os.Stdout).Encode(todos))
		c.HTML(http.StatusOK, "list.html", gin.H{
			"title": "Main website",
			"todos": todos,
		})
	})

	r.GET("/todo/get", func(c *gin.Context) {
		var todo models.Todo
		id, _ := c.GetQuery("id")
		result := dbConnection.First(&todo, id)
		if db.ErrorDB(result, c) {
			return
		}
		fmt.Println(json.NewEncoder(os.Stdout).Encode(todo))
		c.JSON(http.StatusOK, todo)
	})

	r.GET("/todo/delete", func(c *gin.Context) {
		id, _ := c.GetQuery("id")
		result := dbConnection.Delete(&models.Todo{}, id)
		if db.ErrorDB(result, c) {
			return
		}
	})

	r.POST("/todo/update", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.PostForm("id"))
		content := c.PostForm("content")
		var todo models.Todo
		result := dbConnection.Where("id = ?", id).Take(&todo)
		if db.ErrorDB(result, c) {
			return
		}
		todo.Content = content
		result = dbConnection.Save(&todo)
		if db.ErrorDB(result, c) {
			return
		}
	})
}

func main() {
	r := gin.Default()
	dbConnection, err := db.ConnectionDB()

	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	err = dbConnection.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	r.LoadHTMLGlob("templates/*")

	listeners(r, dbConnection)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
