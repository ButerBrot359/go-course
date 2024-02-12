import { Outlet } from "react-router-dom";

	import { PageLayout } from "src/app/widgets/PageLayout";
	
	import { usePopup } from "src/app/shared/lib/hooks/usePopup";
	import { PageHeader } from "src/app/shared/ui/PageHeader/PageHeader";
	
	import { HELLO_WORLD_PAGE_URL_PATH } from "../../consts/urlPaths";
	import { HelloWorldContent } from "./HelloWorldContent";
	
	const HelloWorldPage = () => {
	  // // Uncomment this lines to add Form sidebar
	  // const [isSidebarVisible, openSidebar, closeSidebar] = usePopup(false);
	
	  // const outletContext = {
	  // 	openSidebar,
	  //	closeSidebar,
	  //    backUrl: HELLO_WORLD_PAGE_URL_PATH,
	  // };
	
	  return (
		<PageLayout
		  header={<PageHeader>Тестовый заголовок</PageHeader>}
		  // // Uncomment this lines to add Form sidebar
		  // rightSidebarContent={<Outlet context={outletContext} />}
		  // rightSidebarOpen={isSidebarVisible}
		  // rightSidebarOnClose={closeSidebar}
		  // rightSidebarWidth={540}
		>
		  <HelloWorldContent />
		</PageLayout>
	  );
	};
	
	export default HelloWorldPage;
	