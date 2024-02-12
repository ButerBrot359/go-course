package templates

import (
	"fmt"
	"helpers"
	"model"
)

func CreateAppConfig(_ string, fileData model.FileData) string {
	template := fmt.Sprintf(`import { lazy } from "react";

	import authRoles from "src/app/auth/authRoles";
	
	const [PascalCaseModuleName]App = lazy(() => import("./[PascalCaseModuleName]App"));
	
	export const [PascalCaseModuleName]Config = {
	  auth: authRoles.staffView,
	  settings: {
		layout: {
		  config: {},
		},
	  },
	  routes: [
		{
		  path: "[SnakeCaseModuleName]",
		  element: <[PascalCaseModuleName]App />,
		  children: [
			{
			  path: ":id",
			  element: <></>,
			},
		  ],
		},
	  ],
	};
	`)

	return helpers.ModifyTemplateString(template, fileData)
}
