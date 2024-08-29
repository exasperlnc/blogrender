package blogrender

import (
	"embed"
	"html/template"
	"io"
)

type Post struct {
	Title, Description, Body string
	Tags 										 []string
}

type PostRenderer struct {
	templ *template.Template
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func NewPostRender() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r* PostRenderer) Render(w io.Writer, p Post) error {
	
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}
	return nil
}