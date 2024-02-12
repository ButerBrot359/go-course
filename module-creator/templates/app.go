package templates

import (
	"fmt"
	"helpers"
	"model"
)

type TemplateFunc func(model.FileData) string

func GenerateMainAppPage(pageType string, fileData model.FileData) string {
	generateContent := map[string]TemplateFunc{
		"default": createDefaultApp,
		"tabs":    createTabsApp,
	}

	return generateContent[pageType](fileData)
}

func createDefaultApp(fileData model.FileData) string {
	template := fmt.Sprintf(`import { useEffect } from "react";
	import { useParams } from "react-router-dom";
	
	import FusePageCarded from "@fuse/core/FusePageCarded/FusePageCarded";
	import { useThemeMediaQuery } from "@fuse/hooks";
	
	import { PageHeader } from "src/app/components/PageHeader/PageHeader";
	import { DetailsFormWrapper } from "./modals/DetailsFormWrapper";
	import { [PascalCaseModuleName]Content } from "./[PascalCaseModuleName]Content";
	
	import { useRightModalSidebar } from "src/app/hooks/rightModalSidebar";
	import { isRouteHasId } from "src/app/helpers/check";
	import { MODAL_MODS } from "src/app/constants/mods";
	
	const [PascalCaseModuleName]App = () => {
	  const routeParams = useParams();
	  const { isModalVisible, modalType, closeModal, changeModal } =
		useRightModalSidebar(false, MODAL_MODS.view);
	
	  const isMobile = useThemeMediaQuery((theme) => theme.breakpoints.down("lg"));
	
	  useEffect(() => {
		if (isRouteHasId(routeParams.id)) {
		  changeModal()(MODAL_MODS.view);
		} else {
		  closeModal();
		}
	  }, [routeParams]);
	
	  return (
		<FusePageCarded
		  header={<PageHeader>Тестовая папка</PageHeader>}
		  content={<[PascalCaseModuleName]Content changeModal={changeModal()} />}
		  rightSidebarContent={
			<DetailsFormWrapper
			  closeModal={closeModal}
			  changeModal={changeModal()}
			  modalType={modalType}
			/>
		  }
		  rightSidebarOpen={isModalVisible}
		  rightSidebarWidth={500}
		  scroll={isMobile ? "normal" : "content"}
		/>
	  );
	};
	
	export default [PascalCaseModuleName]App;
	
	`)

	return helpers.ModifyTemplateString(template, fileData)
}

func createTabsApp(fileData model.FileData) string {
	template := `import { useEffect } from "react";
	import { useParams } from "react-router-dom";
	
	import FusePageCarded from "@fuse/core/FusePageCarded/FusePageCarded";
	import { useThemeMediaQuery } from "@fuse/hooks";
	
	import { PageHeader } from "src/app/components/PageHeader/PageHeader";
	import { DetailsFormWrapper } from "./modals/DetailsFormWrapper";
	import { [PascalCaseModuleName]Content } from "./[PascalCaseModuleName]Content";
	
	import { useRightModalSidebar } from "src/app/hooks/rightModalSidebar";
	import { isRouteHasId } from "src/app/helpers/check";
	import { MODAL_MODS } from "src/app/constants/mods";
	
	const [PascalCaseModuleName]App = () => {
	  const routeParams = useParams();
	  const { isModalVisible, modalType, closeModal, changeModal } =
		useRightModalSidebar(false, MODAL_MODS.view);
	
	  const isMobile = useThemeMediaQuery((theme) => theme.breakpoints.down("lg"));
	
	  useEffect(() => {
		if (isRouteHasId(routeParams.id)) {
		  changeModal()(MODAL_MODS.view);
		} else {
		  closeModal();
		}
	  }, [routeParams]);
	
	  return (
		<FusePageCarded
		  header={<PageHeader>Тестовая папка</PageHeader>}
		  content={
			<[PascalCaseModuleName]Content
			  changeModal={changeModal()}
			  closeModal={closeModal}
			/>
		  }
		  rightSidebarContent={
			<DetailsFormWrapper
			  closeModal={closeModal}
			  changeModal={changeModal()}
			  modalType={modalType}
			/>
		  }
		  rightSidebarOpen={isModalVisible}
		  rightSidebarWidth={500}
		  scroll={isMobile ? "normal" : "content"}
		/>
	  );
	};
	
	export default [PascalCaseModuleName]App;
	`

	return helpers.ModifyTemplateString(template, fileData)
}
