package viewhandlers

import (
	"embed"
	"html/template"
	"log"

	"github.com/gin-contrib/multitemplate"
)

type Page struct {
	Template *template.Template
	Layout   string
}

// Add a new composite AddCompositeView
// View is composed from given layouts templates and given viewPath that is the template for main content
// View is registerred with the 'name'
func AddCompositeView(r multitemplate.Renderer, name string, viewPath string, layouts []string, fs embed.FS) multitemplate.Renderer {

	layouts = append(layouts, viewPath)
	tmpl, _ := template.ParseFS(fs, layouts...)
	log.Println("AddCompositeView", name, viewPath, tmpl, layouts)
	r.Add(name, tmpl)
	return r
}

func AddSimpleView(r multitemplate.Renderer, name string, path string, fs embed.FS) multitemplate.Renderer {
	tmpl, _ := template.ParseFS(fs, path)
	r.Add(name, tmpl)
	return r
}

// Get main layout of application.
//
//	______________________
//	|        HEADER      |
//	______________________
//	|      NAVIGATION    |
//	______________________
//	|                    |
//	| CONTENT PLACEHOLDER|
//	|                    |
//	______________________
//	|       FOOTER       |
//	______________________
func GetLayouts() []string {
	layouts := []string{"public/layouts/base.html",
		"public/fragments/head.html",
		"public/fragments/header.html",
		"public/fragments/footer.html",
		"public/fragments/nav.html"}
	return layouts
}

//Get simple layout of application.
//  ----------------------
//  |        HEADER      |
//  ----------------------
//  |                    |
//  | CONTENT PLACEHOLDER|
//  |                    |
//  ----------------------
//  |       FOOTER       |
//  ----------------------

func GetSimpleLayouts() []string {
	layouts := []string{"public/layouts/basesimple.html",
		"public/fragments/head.html",
		"public/fragments/header.html",
		"public/fragments/footer.html"}
	return layouts
}
