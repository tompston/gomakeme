package response

import "fmt"

// Predefined messages for possible situations
const ParamIsNotIntMessage = "Only integers as URL params!"
const PageQueryIsNotIntMessage = "Page query param must be a number!"
const FailedDbConnMessage = "Could not connect to the database!"

// functions that generate the appropriate message for the module.
func FoundOneMessage(module_name string) string {
	return fmt.Sprintf(module_name + " found!")
}

func FoundManyMessage(module_name string) string {
	return fmt.Sprintf(module_name + "s found!")
}

func NotFoundOneMessage(module_name string) string {
	return fmt.Sprintf(module_name + " not found!")
}
func NotFoundManyMessage(module_name string) string {
	return fmt.Sprintf(module_name + "s not found!")
}

func CreatedMessage(module_name string) string {
	return fmt.Sprintf(module_name + " created!")
}

func UpdatedMessage(module_name string) string {
	return fmt.Sprintf(module_name + " updated!")
}

func DeletedMessage(module_name string) string {
	return fmt.Sprintf(module_name + " deleted!")
}