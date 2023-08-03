package viewcon

import (
	"embed"
	"html/template"

	"github.com/gin-contrib/multitemplate"
)

type Page struct {
	Template *template.Template
	Layout   string
}

func AddCompositeTemplate(r multitemplate.Renderer, name string, path string, layouts []string, fs embed.FS) multitemplate.Renderer {
	layouts = append(layouts, path)
	tmpl, _ := template.ParseFS(fs, layouts...)
	r.Add(name, tmpl)
	return r
}

func AddSimpleTemplate(r multitemplate.Renderer, name string, path string, fs embed.FS) multitemplate.Renderer {
	tmpl, _ := template.ParseFS(fs, path)
	r.Add(name, tmpl)
	return r
}

func GetLayouts() []string {
	layouts := []string{"public/layouts/base.html",
		"public/layouts/header.html",
		"public/layouts/footer.html",
		"public/layouts/nav.html"}
	return layouts
}
func GetSimpleLayouts() []string {
	layouts := []string{"public/layouts/basesimple.html",
		"public/layouts/header.html",
		"public/layouts/footer.html"}
	return layouts
}
