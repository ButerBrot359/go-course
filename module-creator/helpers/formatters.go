package helpers

import (
	"model"
	"strings"
)

func ModifyTemplateString(template string, fileData model.FileData) string {
	strWithReplacedModule := strings.ReplaceAll(template, "[PascalCaseModuleName]", fileData.PascalCaseFileName)

	str2WithReplacedModule := strings.ReplaceAll(strWithReplacedModule, "[UppercaseTitle]", fileData.UppercaseTitle)

	finalConfigString := strings.ReplaceAll(str2WithReplacedModule, "[SnakeCaseModuleName]", fileData.SnakeCaseFileName)

	final2ConfigString := strings.ReplaceAll(finalConfigString, "[CamelCaseModuleName]", fileData.CamelCaseFileName)

	return strings.ReplaceAll(final2ConfigString, "[back-quote]", "`")
}
