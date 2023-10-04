package http

import (
	"embed"
	"gosi/coreapi/viewhandlers"

	"github.com/gin-contrib/multitemplate"
)

func loadTemplates(fs embed.FS, r multitemplate.Renderer) multitemplate.Renderer {
	viewhandlers.AddCompositeView(r, "home", "public/home/home.html", viewhandlers.GetLayouts(), fs)
	viewhandlers.AddCompositeView(r, "projects", "public/projects/projects.html", viewhandlers.GetLayouts(), fs)
	viewhandlers.AddCompositeView(r, "admin", "public/admin/admin.html", viewhandlers.GetLayouts(), fs)
	viewhandlers.AddCompositeView(r, "login", "public/auth/login.html", viewhandlers.GetSimpleLayouts(), fs)
	viewhandlers.AddCompositeView(r, "error", "public/error/error.html", viewhandlers.GetSimpleLayouts(), fs)
	return r
}
