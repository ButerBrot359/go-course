package tabs

import (
	"helpers"
	"model"
)

func CreateProxyStyle(fileData model.FileData) string {
	template := `@import "src/app/styles/mixins.scss";

	.first-tab {
	  padding: 15px;
	}
	
	.first-tab__table {
	  @include grid-table-style();
	}
	`

	return helpers.ModifyTemplateString(template, fileData)
}
