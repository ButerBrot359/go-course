package ui

import (
	"fmt"
	"helpers"
	"model"
)

func CreatePage(fileData model.FileData) string {
	template := fmt.Sprintf(`import { Outlet } from "react-router-dom";

	import { PageLayout } from "src/app/widgets/PageLayout";
	
	import { usePopup } from "src/app/shared/lib/hooks/usePopup";
	import { PageHeader } from "src/app/shared/ui/PageHeader/PageHeader";
	
	import { [UppercaseTitle]_PAGE_URL_PATH } from "../../consts/urlPaths";
	import { [PascalCaseModuleName]Content } from "./[PascalCaseModuleName]Content";
	
	const [PascalCaseModuleName]Page = () => {
	  // // Uncomment this lines to add Form sidebar
	  // const [isSidebarVisible, openSidebar, closeSidebar] = usePopup(false);
	
	  // const outletContext = {
	  // 	openSidebar,
	  //	closeSidebar,
	  //    backUrl: [UppercaseTitle]_PAGE_URL_PATH,
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
		  <[PascalCaseModuleName]Content />
		</PageLayout>
	  );
	};
	
	export default [PascalCaseModuleName]Page;
	`)

	return helpers.ModifyTemplateString(template, fileData)
}


func CreatePageContent(fileData model.FileData) string {
	template := fmt.Sprintf(`import { useEffect } from "react";
	import { useLocation, useNavigate } from "react-router-dom";
	
	import { Button } from "@mui/material";
	import Add from "@mui/icons-material/Add";
	import FileDownload from "@mui/icons-material/FileDownload";
	import { toDataSourceRequest } from "@progress/kendo-data-query";
	
	import {
	  DefaultToolbar,
	  GridTable,
	  useDownloadGridExcel,
	  useKendoDataState,
	} from "src/app/entities/GridTable";
	
	import { useFetchData } from "src/app/shared/lib/hooks/useFetchData";
	import { KendoObjToArrModifier } from "src/app/shared/lib/helpers/dataModifiers";
	
	import stl from "./[PascalCaseModuleName]Content.module.scss";
	
	const defaultData = { data: [], total: 0 };
	
	export const [PascalCaseModuleName]Content = () => {
	  const location = useLocation();
	  const navigate = useNavigate();
	
	  const { dataState, dataStateChange } = useKendoDataState();
	  const { excelRef, downloadExcel } = useDownloadGridExcel();
	
	  const [tableData, isDataLoading, fetchTableData] = useFetchData(defaultData);
	  const [allTableData, isAllDataLoading, fetchAllTableData] = useFetchData([]);
	
	  const fetchData = async () => {
		const query = toDataSourceRequest(dataState);
	
		// use your service from api
		fetchTableData(() => {}, { query });
	  };
	
	  const defaultColumnsData = [
		{
		  show: true,
		  field: "field1",
		  title: "field1",
		  width: "250px",
		},
		{
		  show: true,
		  field: "field2",
		  title: "field2",
		  width: "250px",
		},
	  ];
	
	  const leftToolbarConfig = [
		{
		  Element: () => (
			<Button
			  variant="outlined"
			  color="success"
			  startIcon={<FileDownload />}
			  onClick={downloadExcel}
			  disabled={isAllDataLoading}
			>
			  Excel
			</Button>
		  ),
		},
		{
		  Element: () => (
			<Button
			  variant="outlined"
			  color="secondary"
			  startIcon={<Add />}
			  onClick={() => navigate("new")}
			>
			  Добавить
			</Button>
		  ),
		},
	  ];
	
	  useEffect(() => {
		fetchData();
	  }, [dataState]);
	
	  useEffect(() => {
		// use your service from api
		fetchAllTableData(() => {}, {}, [KendoObjToArrModifier]);
	  }, []);
	
	  useEffect(() => {
		if (!location?.state?.refresh) return;
	
		// use your service from api
		fetchAllTableData(() => {}, {}, [KendoObjToArrModifier]);
		fetchData();
	  }, [location]);
	
	  return (
		<div className={stl["[SnakeCaseModuleName]"]}>
		  <GridTable
			sortable
			filterable={false}
			gridData={tableData}
			gridColumnsData={defaultColumnsData}
			gridExcelRef={excelRef}
			gridExcelData={allTableData}
			dataState={dataState}
			dataStateChange={dataStateChange}
			gridOrderLSName={"[SnakeCaseModuleName]__ls-order_table"}
			className={stl["[SnakeCaseModuleName]__table"]}
			onRowDoubleClick={({ dataItem }) => navigate([back-quote]${dataItem.id}[back-quote])}
		  >
			<DefaultToolbar loading={isDataLoading} leftSide={leftToolbarConfig} />
		  </GridTable>
		</div>
	  );
	};
	
	`)

	return helpers.ModifyTemplateString(template, fileData)
}


func CreateContentStyles(fileData model.FileData) string {
	template := fmt.Sprintf(`@import "src/app/styles/mixins.scss";

	.[SnakeCaseModuleName] {
	  padding: 15px;
	}
	
	.[SnakeCaseModuleName]__table {
	  @include grid-table-style();
	
	  height: calc(100vh - 6.4rem - 90px);
	}
	
	`)

	return helpers.ModifyTemplateString(template, fileData)
}