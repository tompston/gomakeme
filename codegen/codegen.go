package codegen

import (
	"fmt"
	"strings"
)

func ConvertToLowecase(input string) string {
	return strings.ToLower(input)
}

func ConvertToLowercasePlural(input string) string {
	return fmt.Sprintf("%s%s", strings.ToLower(input), "s")
}
