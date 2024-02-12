import { useEffect } from "react";
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
	
	import stl from "./HelloWorldContent.module.scss";
	
	const defaultData = { data: [], total: 0 };
	
	export const HelloWorldContent = () => {
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
		<div className={stl["hello-world"]}>
		  <GridTable
			sortable
			filterable={false}
			gridData={tableData}
			gridColumnsData={defaultColumnsData}
			gridExcelRef={excelRef}
			gridExcelData={allTableData}
			dataState={dataState}
			dataStateChange={dataStateChange}
			gridOrderLSName={"hello-world__ls-order_table"}
			className={stl["hello-world__table"]}
			onRowDoubleClick={({ dataItem }) => navigate(`${dataItem.id}`)}
		  >
			<DefaultToolbar loading={isDataLoading} leftSide={leftToolbarConfig} />
		  </GridTable>
		</div>
	  );
	};
	
	