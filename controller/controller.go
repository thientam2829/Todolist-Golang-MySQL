package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"golang/final/config"
	"golang/final/models"

	"github.com/gin-gonic/gin"
)

var (
	id        int
	item      string
	completed int
	view      = template.Must(template.ParseFiles("./views/index.html"))
	database  = config.Database()
)

func Show(c *gin.Context) {
	statement, err := database.Query(`SELECT * FROM todos`)

	if err != nil {
		fmt.Println(err)
	}

	var todos []models.Todo

	for statement.Next() {
		err = statement.Scan(&id, &item, &completed)

		if err != nil {
			fmt.Println(err)
		}

		todo := models.Todo{
			Id:        id,
			Item:      item,
			Completed: completed,
		}

		todos = append(todos, todo)
	}

	data := models.View{
		Todos: todos,
	}

	_ = view.Execute(c.Writer, data)
}

func Add(c *gin.Context) {
	item := c.PostForm("item")

	_, err := database.Exec(`INSERT INTO todos (item) VALUE (?)`, item)

	if err != nil {
		fmt.Println(err)
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	_, err := database.Exec(`DELETE FROM todos WHERE id = ?`, id)

	if err != nil {
		fmt.Println(err)
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}

func Complete(c *gin.Context) {
	id := c.Param("id")

	_, err := database.Exec(`UPDATE todos SET completed = 1 WHERE id = ?`, id)

	if err != nil {
		fmt.Println(err)
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}

func Update(c *gin.Context) {
	newItem := c.PostForm("newItem")
	id := c.Param("id")

	_, err := database.Exec(`UPDATE todos SET item =? WHERE id = ?`, newItem, id)

	if err != nil {
		fmt.Println(err)
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}
