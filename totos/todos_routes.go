package todos

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/todos", index)
	r.GET("/todos/:ID", show)
	r.POST("/todos", store)
	r.PATCH("/todos/:ID", update)
	r.DELETE("/todos/:ID", destroy)

	r.GET("/todos/migrate", func(c *gin.Context) {
		CreateTable()
		c.String(http.StatusOK, "Table Created")
	})
}
