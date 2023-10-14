package template

import (
	"bytes"
	"html/template"
	"io/fs"
	"log"
	"path/filepath"
)

func parseTemplateDirectory(dir string) (*template.Template, error) {
	var paths []string

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			paths = append(paths, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func TemplateSendMail(data map[string]string) bytes.Buffer {
	var body bytes.Buffer
	template, err := parseTemplateDirectory("template/html")
	if err != nil {
		log.Fatalf("failed to parse files template | err : %v", err)
	}

	err = template.ExecuteTemplate(&body, "template-mail-one.gohtml", &data)
	if err != nil {
		log.Fatalf("failed to execute template | err : %v", err)
	}

	return body
}
