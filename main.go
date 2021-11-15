package main

import (
	"embed"
	"fmt"
	"html/template"
	"os"
)

//go:embed www
var fs embed.FS

//go:embed templates/help.template.html
var helpTemplateStr string
var helpTemplate *template.Template

func main() {
	var err error
	helpTemplate, err = template.New("help").Parse(helpTemplateStr)
	if err != nil {
		panic(fmt.Sprintf("error parsing template: %s", err))
	}

	err = startUI()
	if err != nil {
		fmt.Println("error initializing UI:", err)
		os.Exit(1)
	}
}
