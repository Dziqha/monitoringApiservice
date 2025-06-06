package main

import (
	"testtodo/service"

	"gofr.dev/pkg/gofr"
)

type Todo struct {
	ID    int
	Title string
	Done  bool
}

var todos []Todo
var nextID int = 1

func main() {
	app := gofr.New()
	go service.StartMonitor()
	
	app.GET("/todos",func(c *gofr.Context) (any, error) {
		return todos, nil
	})

	app.POST("/todo", func(c *gofr.Context) (any, error) {
		var todo Todo

		if err := c.Bind(&todo); err != nil {
			return nil, err
		}

		todo.ID =nextID
		nextID++
		todos = append(todos, todo)
		return todo, nil
	})

	app.GET("/healty", func(c *gofr.Context) (any, error) {
		return "Monitor is healty", nil
	})
	app.Run()
}