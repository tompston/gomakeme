package codegen

func DbConnForFiber() string {
	return `
	// get db connection
	db, err := database.GetDbConn()
	if err != nil {
		return res.Response(c, 500,nil, err.Error())
	}
	defer db.Close()
	_ = db`
}

func ValidateUrlParamForFiber() string {
	return `

	// validate url param
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return res.Response(c, 400, nil, res.ParamIsNotIntMessage)
	}
	_ = id`
}

func ValidatePayloadForFiber() string {
	return `

	// validate the sent json object
	payload := new(ExampleStruct) // define which struct you want to get
	if err := validate.ValidatePayload(c, payload); err != nil {
		return res.Response(c, 400, nil, err.Error())
	}`

}

func DbConnForGin() string {
	return `
	// get db connection
	db, err := database.GetDbConn()
	if err != nil {
		res.Response(c, 500,nil, res.DbConnErrorMessage(err.Error()))
		return
	}
	defer db.Close()
	_ = db`
}

func ValidateUrlParamForGin() string {
	return `

	// validate url param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.Response(c, 400, nil, res.ParamIsNotIntMessage)
		return
	}
	_ = id`
}

func ValidatePayloadForGin() string {
	return `

	// validate the sent json object
	payload := new(ExampleStruct) // define which struct you want to get
	if err := validate.ValidatePayload(c, payload); err != nil {
		res.Response(c, 400, nil, err.Error())
		return
	}`
}
