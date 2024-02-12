package helpers

import (
	"model"
	"strings"
)

func ModifyTemplateString(template string, fileData model.FileData) string {
	strWithReplacedModule := strings.ReplaceAll(template, "[PascalCaseModuleName]", fileData.PascalCaseFileName)

	finalConfigString := strings.ReplaceAll(strWithReplacedModule, "[SnakeCaseModuleName]", fileData.SnakeCaseFileName)

	return strings.ReplaceAll(finalConfigString, "[back-quote]", "`")
}
