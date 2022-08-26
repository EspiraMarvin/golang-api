package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

// get all todos
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

// add todo
func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)

}

func main() {
	router := gin.Default() // creates server
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.Run("localhost:4000")
}
