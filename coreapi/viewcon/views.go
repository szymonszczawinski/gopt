package viewcon

import "text/template"

type View struct {
	Index Page
	Show  Page
	New   Page
	Edit  Page
}

type Page struct {
	Template *template.Template
	Layout   string
}
