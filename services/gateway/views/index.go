package views

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
