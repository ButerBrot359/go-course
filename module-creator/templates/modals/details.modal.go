package modals

import (
	"fmt"
	"helpers"
	"model"
)

func CreateDetailsModal(fileData model.FileData) string {
	template := fmt.Sprintf(`import { useEffect } from "react";
	import { useNavigate, useParams } from "react-router-dom";
	
	import FuseLoading from "@fuse/core/FuseLoading/FuseLoading";
	
	import { DefaultSidebarWrapper } from "src/app/layouts/SidebarWrappers/DefaultSidebarWrapper";
	import { useFetchData } from "src/app/hooks/fetchData";
	import { DetailsFormProxy } from "./details-form";
	
	import s from "./DetailsFormWrapper.module.scss";
	
	export const DetailsFormWrapper = ({
	  modalType = MODAL_MODS.view,
	  closeModal = () => {},
	  changeModal = () => {},
	}) => {
	  const navigate = useNavigate();
	  const routeParams = useParams();
	
	  const [singleElementData, isDataLoading, fetchSingleElementData] =
		useFetchData();
	
	  const closeSidebarAndSetDefaultRoute = () => {
		closeModal();
	
		navigate("/[SnakeCaseModuleName]");
	  };
	
	  const updateElementData = async (id) => {
		if (!id) return;
	
		try {
		  await fetchSingleElementData(() => console.log("Your service"));
		} catch (error) {
		  closeSidebarAndSetDefaultRoute();
		}
	  };
	
	  const removeElement = async (id) => {
		if (!id) return;
	
		try {
		  await (() => console.log("Service to remove element"))();
	
		  closeSidebarAndSetDefaultRoute();
		} catch (error) {
		  closeSidebarAndSetDefaultRoute();
		}
	  };
	
	  useEffect(() => {
		const { id } = routeParams;
	
		if (id) {
		  updateElementData(id);
		}
	  }, [routeParams]);
	
	  return (
		<DefaultSidebarWrapper
		  modalType={modalType}
		  elementId={routeParams.id}
		  className={s["sidebar-wrapper"]}
		  removeElement={removeElement}
		  closeModal={closeSidebarAndSetDefaultRoute}
		  changeModal={changeModal}
		  isEditable
		>
		  {isDataLoading ? (
			<FuseLoading />
		  ) : (
			<DetailsFormProxy
			  type={modalType}
			  closeModal={closeSidebarAndSetDefaultRoute}
			  data={singleElementData}
			  resetData={singleElementData}
			  restrictionId={routeParams.id}
			  className={
				s["details-form-modal-content"] +
				" " +
				s[[back-quote]details-form-modal-content--${modalType}[back-quote]] 
			  }
			/>
		  )}
		</DefaultSidebarWrapper>
	  );
	};
	`)

	return helpers.ModifyTemplateString(template, fileData)
}
