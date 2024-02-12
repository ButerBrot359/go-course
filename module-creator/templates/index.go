package templates

import (
	"fmt"
	"helpers"
	"model"
)

func CreateIndexFile(fileData model.FileData) string {
	template := fmt.Sprintf(`
	import {
		[UppercaseTitle]_PAGE_URL_PATH,
		[PascalCaseModuleName]MenuTranslate,
	  } from "./consts/urlPaths";
	  import { [PascalCaseModuleName]RoutesConfig } from "./config/[PascalCaseModuleName]RoutesConfig";
	  
	  export {
		[UppercaseTitle]_PAGE_URL_PATH,
		[PascalCaseModuleName]MenuTranslate,
		[PascalCaseModuleName]RoutesConfig,
	  };
	  
	`)

	return helpers.ModifyTemplateString(template, fileData)
}
