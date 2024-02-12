package models

import (
	"fmt"
	"helpers"
	"model"
)

func CreateFormSchema(fileData model.FileData) string {
	template := fmt.Sprintf(`import * as yup from "yup";

	// Add form schema
	export const [CamelCaseModuleName]FormSchema = yup.object().shape({
	
	});
	`)

	return helpers.ModifyTemplateString(template, fileData)
}
