package blogrender

import (
	"html/template"
	"io"
)

type Post struct {
	Title, Description, Body string
	Tags 										 []string
}

const (
	postTemplate = `<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`
)
func Render(w io.Writer, p Post) error {
	tmpl, err := template.New("blog").Parse(postTemplate)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, p); err != nil {
		return err
	}

	return nil
}