package routes

import (
	"todo-app/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine){
	api := r.Group("/api/v1")
	{
		todos := api.Group("/todos")
		{
			todos.GET("", handlers.GetTodos)
			todos.GET("/:id", handlers.GetTodo)
			todos.POST("", handlers.CreateTodo)
            todos.PUT("/:id", handlers.UpdateTodo)
            todos.DELETE("/:id", handlers.DeleteTodo)
		}
	}
}