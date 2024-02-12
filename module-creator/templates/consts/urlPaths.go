package consts

import (
	"fmt"
	"helpers"
	"model"
)

func CreateUrlPaths(fileData model.FileData) string {
	template := fmt.Sprintf(`export const [UppercaseTitle]_PAGE_URL_PATH = "[SnakeCaseModuleName]";

	// Add translation
	export const [PascalCaseModuleName]MenuTranslate = {
	  [PascalCaseModuleName]: "",
	};
	
	`)

	return helpers.ModifyTemplateString(template, fileData)
}
