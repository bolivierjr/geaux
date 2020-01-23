package main

import (
	"fmt"

	t "github.com/bolivierjr/goapi/templates"
)

func main() {
	tmplName := t.Template{Name: "Users"}
	fmt.Println(tmplName.CreateModel())
}
