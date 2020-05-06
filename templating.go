package gutil

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func ParseTemplates(root string) *template.Template {
	templ := template.New("")
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if match, _ := regexp.MatchString(`.*\.(html|tmpl|txt)$`, path); match {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}
		return err
	})

	if err != nil {
		panic(err)
	}

	return templ
}
