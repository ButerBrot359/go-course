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
	
	// // uncomment if has form sidebar
	// const [PascalCaseModuleName]View = lazy(() =>
	//   import("../ui/Modals/[PascalCaseModuleName]View/[PascalCaseModuleName]View")
	// );
	// const [PascalCaseModuleName]Form = lazy(() =>
	//   import("../ui/Modals/[PascalCaseModuleName]Form/[PascalCaseModuleName]Form")
	// );
	
	export const [PascalCaseModuleName]RoutesConfig = {
	  auth: authRoles.userView,
	  routes: [
		{
		  path: [UppercaseTitle]_PAGE_URL_PATH,
		  element: <[PascalCaseModuleName]Page />,
		  exact: true,
		  // // uncomment if has form sidebar
		  // children: [
		  //   {
		  //     path: "new",
		  //     element: <[PascalCaseModuleName]Form />,
		  //   },
		  //   {
		  //     path: ":id",
		  //     element: <[PascalCaseModuleName]View />,
		  //   },
		  //   {
		  //     path: ":id/edit",
		  //     element: <[PascalCaseModuleName]Form />,
		  //   },
		  // ],
		},
	  ],
	};
	
	`)

	return helpers.ModifyTemplateString(template, fileData)
}
