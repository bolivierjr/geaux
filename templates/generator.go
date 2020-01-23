package generator

import (
	"bytes"
	"path/filepath"
	"runtime"
	"text/template"
)

// Get the full path of the template directory
var _, filename, _, _ = runtime.Caller(0)
var directory = filepath.Dir(filename)

// Template struct that generates
// model and controller templates
// wtih given name.
type Template struct {
	Name string
}

// CreateModel will create a new model template
// and stick it in the models directory.
func (t Template) CreateModel() string {
	model := CreateTemplate("model.tmpl", t.Name)
	return model
}

// CreateController will create a new controller
// template and stick it in the controllers directory.
func (t *Template) CreateController() {

}

// CreateTemplate is a helper function that creates a new template
func CreateTemplate(file string, name string) string {
	tmplPath := filepath.Join(directory, file)
	tmpl := template.Must(template.ParseFiles(tmplPath))

	var tmplBytes bytes.Buffer
	err := tmpl.Execute(&tmplBytes, name)

	if err != nil {
		panic(err)
	}

	return tmplBytes.String()
}
