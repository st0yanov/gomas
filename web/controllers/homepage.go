package controllers

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

var tmpl = template.Must(template.ParseGlob(getTemplatesPath()))

// HomepageHandler is responsible for the incoming requests for the homepage
func HomepageHandler(w http.ResponseWriter, r *http.Request) {
	ctx := make(map[string]interface{})

	ctx["homepage"] = "active"

	tmpl.ExecuteTemplate(w, "Homepage", ctx)
}

func getTemplatesPath() string {
	dir, _ := os.Getwd()
	result := ""
	if strings.Contains(dir, "gomasd") {
		result = "../../web/templates/*"
	} else {
		result = "web/templates/*"
	}

	return result
}
