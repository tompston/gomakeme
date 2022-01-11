package main

import (
	"fmt"
	"log"

	"github.com/tompston/gomakeme/codegen"
	"github.com/tompston/gomakeme/input"
)

//go:generate go run main.go

func main() {

	// if the config file exists
	if codegen.PathExists("./gomakeme.yml") {

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

	} else {
		fmt.Println(" ---ERROR---   Could not find gomakeme.yml")
	}
}
