package http

import (
	"embed"
	"html/template"

	"github.com/gin-contrib/multitemplate"
)

func loadTemplates(fs embed.FS) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layouts := getLayouts(fs)
	addCompositeTemplate(r, "home", "public/home/home.html", layouts, fs)
	addCompositeTemplate(r, "projects", "public/projects/projects.html", layouts, fs)
	addCompositeTemplate(r, "admin", "public/admin/admin.html", layouts, fs)
	addCompositeTemplate(r, "login", "public/auth/login.html", getSimpleLayouts(), fs)
	addCompositeTemplate(r, "error", "public/error/error.html", getSimpleLayouts(), fs)
	return r
}

func addCompositeTemplate(r multitemplate.Renderer, name string, path string, layouts []string, fs embed.FS) multitemplate.Renderer {
	layouts = append(layouts, path)
	tmpl, _ := template.ParseFS(fs, layouts...)
	r.Add(name, tmpl)
	return r
}

func addSimpleTemplate(r multitemplate.Renderer, name string, path string, fs embed.FS) multitemplate.Renderer {
	tmpl, _ := template.ParseFS(fs, path)
	r.Add(name, tmpl)
	return r
}

func getLayouts(fs embed.FS) []string {
	layouts := []string{}
	site, err := embed.FS.ReadDir(fs, layoutsDir)
	if err != nil {
		panic(err.Error())
	}
	for _, layout := range site {
		layouts = append(layouts, layoutsDir+"/"+layout.Name())
	}
	return layouts
}
func getSimpleLayouts() []string {
	layouts := []string{"public/layouts/basesimple.html",
		"public/layouts/header.html",
		"public/layouts/footer.html"}
	return layouts
}
