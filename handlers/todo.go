package handlers

import (
	"net/http"
	"todo-app/config"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context){
	var todos []models.Todo
	config.DB.Find(&todos)
	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func CreateTodo(c *gin.Context){
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}	
	config.DB.Create(&input)
	c.JSON(http.StatusCreated, gin.H{"data": input})
}

func GetTodo(c *gin.Context){
	var todo models.Todo
	if err := config.DB.First(&todo, c.Param("id")).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found todo"})
        return
	}
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func UpdateTodo(c *gin.Context){
	var todo models.Todo
	if err := config.DB.First(&todo, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Not found todo"})
        return
    }
	
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	config.DB.Model(&todo).Updates(&input)
    c.JSON(http.StatusOK, gin.H{"data": todo})
}

func DeleteTodo(c *gin.Context) {
    var todo models.Todo
    if err := config.DB.First(&todo, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Not found todo"})
        return
    }

    config.DB.Delete(&todo)
    c.JSON(http.StatusOK, gin.H{"message": "Successfully delete"})
}