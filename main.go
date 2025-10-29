package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID   string `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

var todos = []todo{
	{ID: "1", Task: "Buy groceries", Done: false},
	{ID: "2", Task: "Walk the dog", Done: true},
	{ID: "3", Task: "Read a book", Done: false},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func setTodos(c *gin.Context) {
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func getTodoByID(c *gin.Context) {
	id := c.Param("id")

	for _, todo := range todos {
		if todo.ID == id {
			c.IndentedJSON(http.StatusOK, todo)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func main() {
	router := gin.Default()
	router.GET("/todos-get", getTodos)
	router.POST("/todos-set", setTodos)
	router.DELETE("/todos-delete/:id", deleteTodo)
	router.GET("/todos-get/:id", getTodoByID)
	router.Run("localhost:8080")
}
