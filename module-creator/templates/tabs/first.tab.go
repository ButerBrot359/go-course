package tabs

import (
	"helpers"
	"model"
)

func CreateFirstTab(fileData model.FileData) string {
	template := `import { useEffect } from "react";
	import { useNavigate, useParams } from "react-router-dom";
	
	import { toDataSourceRequest } from "@progress/kendo-data-query";
	import { FileDownload } from "@mui/icons-material";
	import AddIcon from "@mui/icons-material/Add";
	
	import { DefaultToolbar } from "src/app/layouts/KendoUIToolbars/DefaultToolbar";
	import { GridDynamicTable } from "src/app/layouts/KendoUIGrid/GridDynamicTable";
	import { SuccessButton } from "src/app/components/Buttons/SuccessButton";
	
	import { useKendoDataState } from "src/app/hooks/kendo-data-state";
	import { useFetchData } from "src/app/hooks/fetchData";
	import { useDownload } from "src/app/hooks/download";
	
	import { DateFormat } from "src/app/helpers/formatters";
	import { MODAL_MODS } from "src/app/constants/mods";
	
	import s from "./index.module.scss";
	
	export const FirstTab = ({ changeModal = () => {} }) => {
	  const navigate = useNavigate();
	  const routeParams = useParams();
	
	  const { dataState, dataStateChange } = useKendoDataState();
	  const { downloadRef, download } = useDownload();
	
	  const [elementsData, isDataLoading, fetchElementsData] = useFetchData();
	  const [allElementsData, isAllDataLoading, fetchAllElementsData] =
		useFetchData();
	
	  const defaultColumnsData = [
		{
		  orderIndex: 0,
		  show: true,
		  field: "",
		  title: "Номер-тест",
		  width: "auto",
		  filterable: false,
		},
		{
		  orderIndex: 1,
		  show: true,
		  field: "",
		  title: "Клиент-тест",
		  width: "auto",
		  filterable: false,
		},
		{
		  orderIndex: 2,
		  show: true,
		  field: "",
		  title: "Сумма-тест",
		  width: "auto",
		  filterable: false,
		},
		{
		  orderIndex: 3,
		  show: true,
		  field: "",
		  title: "Дата",
		  width: "auto",
		  cell: (props) => {
			const date = new Date(); // pass your dates field instead new Date()
			return <td>{[back-quote]${date ? DateFormat(date) : "-"}[back-quote]}</td>;
		  },
		},
	  ];
	
	  const leftToolbarConfig = [
		{
		  Element: () => (
			<SuccessButton
			  startIcon={<FileDownload />}
			  onClick={download}
			  disabled={isAllDataLoading}
			>
			  Excel
			</SuccessButton>
		  ),
		},
		{
		  Element: () => (
			<SuccessButton
			  startIcon={<AddIcon />}
			  onClick={() => changeModal(MODAL_MODS.create)}
			>
			  Добавить
			</SuccessButton>
		  ),
		},
	  ];
	
	  const openSingleElement = ({ dataItem }) => {
		const { id } = dataItem;
	
		navigate(id);
	  };
	
	  useEffect(() => {
		const query = toDataSourceRequest(dataState);
	
		fetchElementsData(() => console.log("Your service"), { query });
	  }, [dataState, routeParams]);
	
	  useEffect(() => {
		fetchAllElementsData(() => console.log("Your service"));
	  }, [routeParams]);
	
	  return (
		<div className={s["first-tab"]}>
		  <GridDynamicTable
			filterable
			sortable
			gridData={elementsData}
			gridColumnsData={defaultColumnsData}
			gridExcelRef={downloadRef}
			gridExcelData={allElementsData?.data || []}
			dataState={dataState}
			dataStateChange={dataStateChange}
			gridOrderLSName={"test-folder-table"}
			className={s["first-tab__table"]}
			onRowDoubleClick={openSingleElement}
		  >
			<DefaultToolbar loading={isDataLoading} leftSide={leftToolbarConfig} />
		  </GridDynamicTable>
		</div>
	  );
	};
	`
return helpers.ModifyTemplateString(template, fileData)
}
