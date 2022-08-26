package main

import (
	"errors"
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

// display home page
func HomepageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Todo APIs with Golang"})
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

// get specific todo
func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	// if todo doesn't exist we throw an error
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo Not Found"})
	}

	context.IndentedJSON(http.StatusOK, todo)

}

// get todo by its id
func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

// update todo
func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	// if todo doesn't exist we throw an error
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo Not Found"})
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func deleteTodo(context *gin.Context) {
	id := context.Param("id")
	index := -1
	// todo, _ := getTodoById(id)
	// if err != nil {
	// 	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo Not Found"})
	// }

	for i := 0; i < len(todos); i++ {
		if (todos[i].ID) == id {
			index = 1
		}
	}

	if index == -1 {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Todo not found",
		})
		return
	}

	todos = append(todos[:index], todos[index+1:]...)
	// context.JSON(http.StatusOK, todo)
	context.JSON(http.StatusOK, gin.H{"message": "Todo has been deleted"})
	// context.IndentedJSON(http.StatusOK, todo, gin.H{ "message": "Todo has been deleted"})
}

func main() {
	router := gin.Default() // creates server
	router.GET("/", HomepageHandler)
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.DELETE("/todos/:id", deleteTodo)
	router.Run("localhost:4000")
}
