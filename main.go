package main

import (
	"gomakeme/codegen"
	"gomakeme/input"
	"log"
)

//go:generate go run main.go

func main() {

	conf, err := input.NewConfig("gomakeme.yml")
	if err != nil {
		log.Fatal(err)
	}

	// generate init project
	codegen.GenerateInitProject(*conf)

	// generate the modules specified in the .yml file
	codegen.GenerateModulesFromConfig(conf)

	// keep the modules specified in the yaml file in sync with the modules that are imported to the router
	codegen.UpdateModuleImporter(*conf)

}
