package todos

import (
	"net/http"

	"danielweaver.dev/go-todo/utils"
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	todos := GetAllTodos()
	c.HTML(http.StatusOK, "todos_index.html", gin.H{
		"todos":   todos,
		"message": "",
	})
}

func show(c *gin.Context) {
	// TODO: WIP
}

func store(c *gin.Context) {
	var todo Todo
	err := c.ShouldBind(&todo)
	utils.CheckError(err, "Error parsing request to store todo.")
	StoreTodo(&todo)

	renderAllTodos(c, "Saved!")
}

func update(c *gin.Context) {
	var todo Todo
	err := c.ShouldBind(&todo)
	utils.CheckError(err, "Error parsing request to update toto.")
	UpdateTodo(&todo)

	renderAllTodos(c, "Updated!")
}

func destroy(c *gin.Context) {
	var todo Todo
	err := c.ShouldBindUri(&todo)
	utils.CheckError(err, "Error parsing request to delete toto.")
	DeleteTodo(&todo)

	renderAllTodos(c, "Deleted!")
}

func renderAllTodos(c *gin.Context, message string) {
	todos := GetAllTodos()
	c.HTML(http.StatusOK, "todos_list.html", gin.H{
		"todos":   todos,
		"message": message,
	})
}
