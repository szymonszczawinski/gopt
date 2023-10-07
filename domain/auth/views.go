package auth

import "html/template"

type Page struct {
	name     string
	template *template.Template
}
