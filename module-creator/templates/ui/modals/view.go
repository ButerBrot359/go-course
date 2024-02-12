package modals

import (
	"fmt"
	"helpers"
	"model"
)

func CreateModalView(fileData model.FileData) string {
	template := fmt.Sprintf(`import { useEffect } from "react";
	import { useOutletContext, useParams } from "react-router-dom";
	
	import FuseLoading from "@fuse/core/FuseLoading";
	import { Divider } from "@mui/material";
	
	import { ViewHeader } from "src/app/entities/FormViewWrappers";
	
	import { useFetchData } from "src/app/shared/lib/hooks/useFetchData";
	import { FormViewRow } from "src/app/shared/ui/FormViewRow/FormViewRow";
	import { useNavigateToRoute } from "src/app/shared/lib/hooks/useNavigateToRoute";
	
	import stl from "./[PascalCaseModuleName]View.module.scss";
	
	const [PascalCaseModuleName]View = () => {
	  const { id } = useParams();
	  const [selectedElement, isSelectedElementLoading, fetchSelectedElement] =
		useFetchData({});
	
	  const { closeSidebar, openSidebar, backUrl } = useOutletContext();
	
	  const [navigateBack] = useNavigateToRoute(backUrl);
	
	  const deleteElement = async () => {
		// remove element service
		await new Promise((res) => res())({ id });
	
		navigateBack({ refresh: true })();
	  };
	
	  useEffect(() => {
		openSidebar();
	
		// get single element service
		fetchSelectedElement(new Promise((res) => res()), { id });
	
		return () => {
		  closeSidebar();
		};
	  }, [id]);
	
	  if (isSelectedElementLoading) {
		return (
		  <div className={stl["loading-wrapper"]}>
			<FuseLoading />
		  </div>
		);
	  }
	
	  const { test } = selectedElement;
	
	  return (
		<div className={stl["[SnakeCaseModuleName]-view"]}>
		  <ViewHeader
			closeAction={navigateBack({ refresh: false })}
			deleteAction={deleteElement}
		  />
		  <main>
			<Divider />
			<FormViewRow title="test1">{test || "-"}</FormViewRow>
		  </main>
		</div>
	  );
	};
	
	export default [PascalCaseModuleName]View;
	`)

	return helpers.ModifyTemplateString(template, fileData)
}

func CreateModalViewStyle(fileData model.FileData) string {
	template := fmt.Sprintf(`@import "src/app/styles/mixins.scss";

	.[SnakeCaseModuleName]-view {
	  @include sidebar-form-wrapper();
	}
	
	`)

	return helpers.ModifyTemplateString(template, fileData)
}
