package templates

import (
	"fmt"
	"helpers"
	"model"
)

func GenerateContentPage(pageType string, fileData model.FileData) string {
	generateContent := map[string]TemplateFunc{
		"default": CreateDefautContent,
		"tabs":    CreateTabsContent,
	}

	return generateContent[pageType](fileData)
}

func CreateDefautContent(fileData model.FileData) string {
	template := fmt.Sprintf(`import { useEffect } from "react";
	import { useNavigate, useParams } from "react-router-dom";
	
	import { toDataSourceRequest } from "@progress/kendo-data-query";
	import { FileDownload } from "@mui/icons-material";
	import AddIcon from "@mui/icons-material/Add";
	
	import { GridDynamicTable } from "src/app/layouts/KendoUIGrid/GridDynamicTable";
	import { DefaultToolbar } from "src/app/layouts/KendoUIToolbars/DefaultToolbar";
	import { SuccessButton } from "src/app/components/Buttons/SuccessButton";
	
	import { useKendoDataState } from "src/app/hooks/kendo-data-state";
	import { useFetchData } from "src/app/hooks/fetchData";
	import { useDownload } from "src/app/hooks/download";
	
	import { MODAL_MODS } from "src/app/constants/mods";
	
	import s from "./[PascalCaseModuleName].module.scss";
	
	export const [PascalCaseModuleName]Content = ({ changeModal = () => {} }) => {
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
		<div className={s["[SnakeCaseModuleName]"]}>
		  <GridDynamicTable
			filterable
			sortable
			gridData={elementsData}
			gridColumnsData={defaultColumnsData}
			gridExcelRef={downloadRef}
			gridExcelData={allElementsData?.data || []}
			dataState={dataState}
			dataStateChange={dataStateChange}
			gridOrderLSName={"[SnakeCaseModuleName]-table"}
			className={s["[SnakeCaseModuleName]__table"]}
			onRowDoubleClick={openSingleElement}
		  >
			<DefaultToolbar loading={isDataLoading} leftSide={leftToolbarConfig} />
		  </GridDynamicTable>
		</div>
	  );
	};
	`)

	return helpers.ModifyTemplateString(template, fileData)
}

func CreateTabsContent(fileData model.FileData) string {
	template := `import { TabsWrapper } from "src/app/layouts/MUIExtensions/TabsWrapper";
	import { useTabsSwitcher } from "src/app/hooks/tabsSwitcher";
	import { [PascalCaseModuleName]TabsProxy } from "./tabs";
	
	const contentTabsName = [
	  { type: "firstTab", label: "Первый таб" },
	  { type: "secondTab", label: "Второй таб" },
	];
	
	export const [PascalCaseModuleName]Content = ({
	  changeModal = () => {},
	  closeModal = () => {},
	}) => {
	  const { currentTabIndex, handleTabChange } = useTabsSwitcher();
	
	  const changeTab = (_, value) => {
		handleTabChange(_, value);
		closeModal();
	  };
	
	  return (
		<TabsWrapper
		  currentTabIndex={currentTabIndex}
		  handleTabChange={changeTab}
		  tabsContent={contentTabsName}
		>
		  <[PascalCaseModuleName]TabsProxy
			type={contentTabsName[currentTabIndex].type}
			changeModal={changeModal}
		  />
		</TabsWrapper>
	  );
	};
	`

	return helpers.ModifyTemplateString(template, fileData)
}
