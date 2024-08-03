package views

import (
	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/to-rename/internal/scrapModels"
)

func (h *Handler) index(c *gin.Context) {
	data := scrapModels.MovieFull{
		Name:        "Ass",
		NameEng:     "Eng Ass",
		Description: "Someting stupid",
	}

	c.HTML(200, "index.html", data)
}
