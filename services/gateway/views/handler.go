package views

import (
	"embed"
	"html/template"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*.html
var TS embed.FS

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HandleAll() {
	r := gin.Default()
	tmpl := template.Must(template.ParseFS(TS, "templates/*.html"))
	r.SetHTMLTemplate(tmpl)

	// future routes here
	r.GET("/index", h.index)

	r.Run(":8080")
}
