package web

import (
	"html/template"
	"io"
	"io/fs"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer implements echo.Renderer (Echo has no built-in implementation).
// Templates are parsed once at startup and cached for the lifetime of the process.
type TemplateRenderer struct {
	templates *template.Template
}

// NewTemplateRenderer recursively walks templatesDir, parses every *.html file
// found (layout, pages, partials) into a single shared template set, and
// returns a renderer ready to be assigned to echo.Echo.Renderer.
func NewTemplateRenderer(templatesDir string) *TemplateRenderer {
	var files []string

	err := filepath.WalkDir(templatesDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ".html" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	return &TemplateRenderer{
		templates: template.Must(template.ParseFiles(files...)),
	}
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
