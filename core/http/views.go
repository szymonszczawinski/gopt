package http

import (
	"embed"
	"gosi/coreapi/viewcon"

	"github.com/gin-contrib/multitemplate"
)

func loadTemplates(fs embed.FS, r multitemplate.Renderer) multitemplate.Renderer {
	viewcon.AddCompositeTemplate(r, "home", "public/home/home.html", viewcon.GetLayouts(), fs)
	viewcon.AddCompositeTemplate(r, "projects", "public/projects/projects.html", viewcon.GetLayouts(), fs)
	viewcon.AddCompositeTemplate(r, "admin", "public/admin/admin.html", viewcon.GetLayouts(), fs)
	viewcon.AddCompositeTemplate(r, "login", "public/auth/login.html", viewcon.GetSimpleLayouts(), fs)
	viewcon.AddCompositeTemplate(r, "error", "public/error/error.html", viewcon.GetSimpleLayouts(), fs)
	return r
}
