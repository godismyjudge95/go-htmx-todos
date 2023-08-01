package main

import (
	todos "danielweaver.dev/go-todo/totos"
	"danielweaver.dev/go-todo/utils"
	"github.com/bmatcuk/doublestar/v4"
	"github.com/gin-gonic/gin"
)

func setup() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	templates, err := doublestar.FilepathGlob("*/templates/**/*.html")
	utils.CheckError(err, "Error attempting to load template files.")
	r.LoadHTMLFiles(templates...)
	todos.Routes(r)

	return r
}

func main() {
	r := setup()
	r.Run(":8080")
}
