package tabs

import (
	"helpers"
	"model"
)

func CreateSecondTab(fileData model.FileData) string {
	template := `export const SecondTab = () => {
		return <></>;
	  };
	  `

	return helpers.ModifyTemplateString(template, fileData)
}
