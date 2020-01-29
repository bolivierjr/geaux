package templates

// HandlerTemplate struct that generates
// a handler template wtih given name.
type HandlerTemplate struct {
	Name string
}

// NewHandler initializes a HandlerTemplate.
func NewHandler(name string) HandlerTemplate {
	return HandlerTemplate{Name: name}
}

// CreateHandler will create a new handler
// template and stick it in the handlers directory.
func (t HandlerTemplate) CreateHandler() {

}
