package routes

import (
	"journal/controllers"
	"journal/middleware"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(router *gin.Engine) {
    router.Use(middleware.Auth)
    
    todoGroup := router.Group("/todos")
    {
        todoGroup.GET("/", controllers.GetTodoList)
        todoGroup.POST("/", controllers.CreateTodo)
        todoGroup.PUT("/:id", controllers.UpdateTodo)
    }
}
