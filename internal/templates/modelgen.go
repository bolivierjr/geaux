package templates

import (
	"bytes"
	"path/filepath"
	"runtime"
	"text/template"
)

// Get the full path of the template directory.
var (
	_, filename, _, _ = runtime.Caller(0)
	directory         = filepath.Dir(filename)
)

// ModelTemplate struct that generates
// a model templates wtih given name.
type ModelTemplate struct {
	Name string
}

// NewModel initializes a ModelTemplate.
func NewModel(name string) ModelTemplate {
	return ModelTemplate{Name: name}
}

// CreateModel will create a new model template
// and stick it in the models directory.
func (t ModelTemplate) CreateModel() string {
	model := CreateTemplate("model.tmpl", t.Name)
	return model
}

// CreateTemplate is a helper function that creates a new template.
func CreateTemplate(file string, name string) string {
	tmplFile := filepath.Join(directory, file)
	tmpl := template.Must(template.ParseFiles(tmplFile))

	var tmplBytes bytes.Buffer
	err := tmpl.Execute(&tmplBytes, name)
	if err != nil {
		panic(err)
	}

	return tmplBytes.String()
}
