package templates

import (
	"fmt"
	"helpers"
	"model"
)

func GenerateAppStyle(pageType string, fileData model.FileData) string {
	generateContent := map[string]TemplateFunc{
		"default": createDefaultAppStyle,
		"tabs":    createTabsAppStyle,
	}

	return generateContent[pageType](fileData)
}

func createDefaultAppStyle(fileData model.FileData) string {
	template := fmt.Sprintf(`@import "src/app/styles/mixins.scss";

	.[SnakeCaseModuleName] {
	  padding: 15px;
	}
	
	.[SnakeCaseModuleName]__table {
	  @include grid-table-style();
	}`)

	return helpers.ModifyTemplateString(template, fileData)
}

func createTabsAppStyle(fileData model.FileData) string {
	template := fmt.Sprintf(``)

	return helpers.ModifyTemplateString(template, fileData)
}
