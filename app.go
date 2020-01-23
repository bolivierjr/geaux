package main

import (
	"fmt"
    m "github.com/bolivierjr/goapi/templates"
)


type Template struct {
	Name string
}

func main() {
	tmplName := m.Template{"Users"}
	fmt.Println(tmplName.CreateModel())
}
