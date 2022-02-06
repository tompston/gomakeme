package codegen

import (
	"embed"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

// -- init embeds
//go:embed templates/init_server/*
//go:embed templates/init_server/*/*
//go:embed templates/init_server/*/*/*

// -- module embeds
//go:embed templates/init_module/*

// -- optional embeds
//go:embed templates/optional/*
//go:embed templates/optional/*/*

var templates embed.FS

// var debug_mode = true
var debug_mode = false

// pass in a locally defined string + template functions and return it as a template
// # template_variables, err := CreateLocalTemplate(module_variables, temp_funcs)
func CreateLocalTemplate(input string, t_funcs template.FuncMap) (*template.Template, error) {

	tpl, err := template.New("test").Funcs(t_funcs).Parse(input)

	if err != nil {
		return nil, err
	}
	return tpl, err
}

func ExecuteTemplate(t_name string, t_path string, t_funcs template.FuncMap, full_output_path string, data interface{}) {

	tmpl, err := template.New(t_name).Funcs(t_funcs).ParseFS(templates, t_path)

	out, _ := os.Create(full_output_path)

	defer out.Close()

	if err != nil {
		log.Fatalf("Could not parse struct template: %v\n", err)
	}

	tmpl.Execute(out, data)
}

// generate 2 strings from the template path, that are needed to run the ExecuteTemplate() func
//    if return_full_path = false, return only the .go file from the input string
//    if return_full_path = true, return the full input string minus the .tpl extension
// The boolean option is used so that you could use the full_path string to generate a replica
// of the `init_server` directory easier
func GenerateTemplateNameAndOutput(template_path string, return_full_path bool) (string, string) {

	// return everything that is after the last / to get the name of the template.
	temp_name := regexp.MustCompile(`[^/]*$`).FindString(template_path)

	if !return_full_path {
		out := strings.TrimSuffix(temp_name, filepath.Ext(temp_name))
		return temp_name, out
	}

	out := strings.TrimSuffix(template_path, filepath.Ext(template_path))

	return temp_name, out
}
