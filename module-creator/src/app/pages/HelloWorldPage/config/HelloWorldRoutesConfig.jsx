import { lazy } from "react";
	import authRoles from "src/app/auth/authRoles";
	
	import { HELLO_WORLD_PAGE_URL_PATH } from "../consts/urlPaths";
	
	const HelloWorldPage = lazy(() =>
	  import("../ui/HelloWorldPage/HelloWorldPage")
	);
	
	// // uncomment if has form sidebar
	// const HelloWorldView = lazy(() =>
	//   import("../ui/Modals/HelloWorldView/HelloWorldView")
	// );
	// const HelloWorldForm = lazy(() =>
	//   import("../ui/Modals/HelloWorldForm/HelloWorldForm")
	// );
	
	export const HelloWorldRoutesConfig = {
	  auth: authRoles.userView,
	  routes: [
		{
		  path: HELLO_WORLD_PAGE_URL_PATH,
		  element: <HelloWorldPage />,
		  exact: true,
		  // // uncomment if has form sidebar
		  // children: [
		  //   {
		  //     path: "new",
		  //     element: <HelloWorldForm />,
		  //   },
		  //   {
		  //     path: ":id",
		  //     element: <HelloWorldView />,
		  //   },
		  //   {
		  //     path: ":id/edit",
		  //     element: <HelloWorldForm />,
		  //   },
		  // ],
		},
	  ],
	};
	
	