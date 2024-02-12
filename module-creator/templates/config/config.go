package fileconfig

import (
	"fmt"
	"helpers"
	"model"
)

func CreateConfigFile(fileData model.FileData) string {
	template := fmt.Sprintf(`import { lazy } from "react";
	import authRoles from "src/app/auth/authRoles";
	
	import { [UppercaseTitle]_PAGE_URL_PATH } from "../consts/urlPaths";
	
	const [PascalCaseModuleName]Page = lazy(() =>
	  import("../ui/[PascalCaseModuleName]Page/[PascalCaseModuleName]Page")
	);
	
	export const [PascalCaseModuleName]RoutesConfig = {
	  auth: authRoles.userView,
	  routes: [
		{
		  path: [UppercaseTitle]_PAGE_URL_PATH,
		  element: <[PascalCaseModuleName]Page />,
		  exact: true,
		},
	  ],
	};
	
	`)

	return helpers.ModifyTemplateString(template, fileData)
}
