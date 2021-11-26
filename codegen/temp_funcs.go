package codegen

import (
	"strings"
	"text/template"
)

// functions that will be passed to the templates.
var temp_funcs = template.FuncMap{
	"convertToTitle":     strings.Title,
	"convertToLowercase": strings.ToLower,
	"db_conn_snippet":    DbConnSnippet,
	"validate_url_param": ValidateUrlParam,
	"validate_payload":   ValidatePayload,
}

// --- Funcs that return the code snippet that can be used in the template

func DbConnSnippet() string {
	return `
	// get db connection
	db, err := database.GetDbConn()
	if err != nil {
		return response.ResponseError(c, nil, err.Error())
	}
	defer db.Close()
	_ = db`
}

func ValidateUrlParam() string {
	return `

	// validate url param
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.ResponseError(c, nil, response.ParamIsNotIntMessage)
	}
	_ = id`
}

func ValidatePayload() string {
	return `

	// validate the sent json object
	payload := new(ExampleStruct) // define which struct you want to get
	if err := validate.ValidatePayload(c, payload); err != nil {
		return response.ResponseError(c, nil, err.Error())
	}`
}
