package routes

import (
	"go-todolist/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine) {
	router := gin.Default()

	router.GET("/list", controllers.GetTodoList)
}
