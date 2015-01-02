package config

import (
	"html/template"
	"log"
	"io"
)

func Render(wr io.Writer, template *template.Template, data interface{}) {
	err := template.Execute(wr, data)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}