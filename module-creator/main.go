package main

import (
	"bufio"
	"consts"
	"fileconfig"
	"flag"
	"fmt"
	"log"
	"model"
	"models"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	moduleName := flag.String("moduleName", "", "Module name")
	hasTabs := flag.Bool("t", false, "Add tabs folder")

	flag.Parse()

	currentDirectory, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		return
	}

	if *moduleName == "" {
		fmt.Println("Empty module name")
		return
	}

	var newFileData = &model.FileData{FolderName: *moduleName, HasTabs: *hasTabs}

	newFileData.SetPascalCase()
	newFileData.SetSnakeCase()
	newFileData.SetCamelCase()
	newFileData.SetUppercaseTitle()
	newFileData.SetRootDirectory(currentDirectory)

	fmt.Println(newFileData)

	pageModuleName := fmt.Sprintf(`%vPage`, newFileData.PascalCaseFileName);
	moduleDirPath := filepath.Join(newFileData.RootDirectory,"src", "app", "pages", pageModuleName)

	// configDirPath := filepath.Join(newFileData.RootDirectory, "src", "app", "configs")

	// if err := rewriteConfigFiles(configDirPath, newFileData); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	if err := createDir(moduleDirPath); err != nil {
		log.Fatal("Folder already exist")
		return
	}

	folders := []string{"api", "config", "consts", "models", "ui", "utils"}
	
	generateFoldersInPage(moduleDirPath, folders)
	generateConfigContent(filepath.Join(moduleDirPath, "config"), newFileData)
	generateConstsContent(filepath.Join(moduleDirPath, "consts"), newFileData)
	generateFormSchemaContent(filepath.Join(moduleDirPath, "models"), newFileData)

	// generateMainContent(newFileData, moduleDirPath)
	// generateDetailsModal(newFileData, moduleDirPath)

	// if newFileData.HasTabs {
	// 	generateTabsContent(newFileData, moduleDirPath)
	// }

}

// func generateMainContent(newFileData *model.FileData, moduleDirPath string) {
// 	rootFiles := [4]model.FileProperty{
// 		{Postfix: "App", Extensions: []string{"jsx"}, BuilderFunc: templates.GenerateMainAppPage},
// 		{Postfix: "Content", Extensions: []string{"jsx"}, BuilderFunc: templates.GenerateContentPage},
// 		{Postfix: "Config", Extensions: []string{"js"}, BuilderFunc: templates.CreateAppConfig},
// 		{Postfix: "", Extensions: []string{"module", "scss"}, BuilderFunc: templates.GenerateAppStyle},
// 	}

// 	for _, fileProps := range rootFiles {
// 		fullFileName := model.AddExtensions(newFileData.AddPostfix(fileProps.Postfix), fileProps.Extensions...)

// 		filePath := filepath.Join(moduleDirPath, fullFileName)

// 		var fileConent string

// 		if newFileData.HasTabs {
// 			fileConent = fileProps.BuilderFunc(model.TabsFileType, *newFileData)
// 		} else {
// 			fileConent = fileProps.BuilderFunc(model.DefaultFileType, *newFileData)
// 		}


// 		// createFileAndWriteData(filePath, fileConent)
// 	}
// }

// func generateDetailsModal(newFileData *model.FileData, moduleDirPath string) {
// 	modalDir := filepath.Join(moduleDirPath, "modals")
// 	if err := createDir(modalDir); err != nil {
// 		log.Fatal("Folder already exist")
// 		return
// 	}

// 	rootFiles := [2]model.CustomFileProperty{
// 		{Name: "DetailsFormWrapper", Extensions: []string{"jsx"}, BuilderFunc: modals.CreateDetailsModal},
// 		{Name: "DetailsFormWrapper", Extensions: []string{"module", "scss"}, BuilderFunc: modals.CreateDetailsStyle},
// 	}

// 	for _, fileProps := range rootFiles {
// 		fullFileName := model.AddExtensions(fileProps.Name, fileProps.Extensions...)

// 		filePath := filepath.Join(modalDir, fullFileName)

// 		fileConent := fileProps.BuilderFunc(*newFileData)

// 		createFileAndWriteData(filePath, fileConent)
// 	}

// 	createDetailsProxyfolder(newFileData, modalDir)
// }

// func createDetailsProxyfolder(newFileData *model.FileData, moduleDirPath string) {
// 	modalDir := filepath.Join(moduleDirPath, "details-form")
// 	if err := createDir(modalDir); err != nil {
// 		log.Fatal("Folder already exist")
// 		return
// 	}

// 	rootFiles := []model.CustomFileProperty{
// 		{Name: "index", Extensions: []string{"jsx"}, BuilderFunc: details.CreateDetailProxy},
// 		{Name: "index", Extensions: []string{"module", "scss"}, BuilderFunc: details.CreateDetailProxyStyle},
// 		{Name: "DetailsFormModel", Extensions: []string{"js"}, BuilderFunc: details.CreateDetailModel},
// 		{Name: "DetailsViewForm", Extensions: []string{"jsx"}, BuilderFunc: details.CreateDetailView},
// 		{Name: "DetailsCreateForm", Extensions: []string{"jsx"}, BuilderFunc: details.CreateDetailCreate},
// 		{Name: "DetailsEditForm", Extensions: []string{"jsx"}, BuilderFunc: details.CreateDetailEdit},
// 	}

// 	for _, fileProps := range rootFiles {
// 		fullFileName := model.AddExtensions(fileProps.Name, fileProps.Extensions...)

// 		filePath := filepath.Join(modalDir, fullFileName)

// 		fileConent := fileProps.BuilderFunc(*newFileData)

// 		createFileAndWriteData(filePath, fileConent)
// 	}
// }

// func generateTabsContent(newFileData *model.FileData, moduleDirPath string) {
// 	tabsDir := filepath.Join(moduleDirPath, "tabs")
// 	if err := createDir(tabsDir); err != nil {
// 		log.Fatal("Folder already exist")
// 		return
// 	}

// 	rootFiles := []model.CustomFileProperty{
// 		{Name: "index", Extensions: []string{"jsx"}, BuilderFunc: tabs.CreateProxyTab},
// 		{Name: "index", Extensions: []string{"module", "scss"}, BuilderFunc: tabs.CreateProxyStyle},
// 		{Name: "FirstTab", Extensions: []string{"jsx"}, BuilderFunc: tabs.CreateFirstTab},
// 		{Name: "SecondTab", Extensions: []string{"jsx"}, BuilderFunc: tabs.CreateSecondTab},
// 	}

// 	for _, fileProps := range rootFiles {
// 		fullFileName := model.AddExtensions(fileProps.Name, fileProps.Extensions...)

// 		filePath := filepath.Join(tabsDir, fullFileName)

// 		fileConent := fileProps.BuilderFunc(*newFileData)

// 		createFileAndWriteData(filePath, fileConent)
// 	}
// }

func rewriteConfigFiles(configDirPath string, newFileData *model.FileData) error {
	urlConfigFilePath := filepath.Join(configDirPath, "urlsConfig.js");
	routesConfigFilePath := filepath.Join(configDirPath, "routesConfig.js");
	translateConfigFilePath := filepath.Join(configDirPath, "navigation-i18n", "en.js")

	if err := rewriteUrlConfig(urlConfigFilePath, newFileData); err != nil {
		return err
	}

	if err := rewriteRoutesConfig(routesConfigFilePath, newFileData); err != nil {
		return err
	}

	if err := rewriteTranslateConfig(translateConfigFilePath,newFileData); err != nil {
		return err
	}
	
	return nil
}

func rewriteUrlConfig(filePath string, newFileData *model.FileData) error {
	pageURLPath := fmt.Sprintf(`%v_PAGE_URL_PATH`, newFileData.UppercaseTitle);
	importLine := fmt.Sprintf(`import { %v } from "../pages/%vPage";`, pageURLPath, newFileData.PascalCaseFileName);


	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
    if err != nil {
        fmt.Println(err)
        return nil
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
    var lines []string
	lines = append(lines, fmt.Sprintf(`%v`, importLine) + "\n")

	isInsideUrlArr := false;

    for scanner.Scan() {
		if strings.Contains(scanner.Text(), pageURLPath) {
			return fmt.Errorf("Module with this URL config already exist")
		}

		if strings.Contains(scanner.Text(), "urls") {
			isInsideUrlArr = true
		}

		if strings.Contains(scanner.Text(), "];") && isInsideUrlArr {
			lines = append(lines, fmt.Sprintf("\t" + `%v`, pageURLPath) + "\n")
			isInsideUrlArr = false
		}

        lines = append(lines, scanner.Text())
    }

    // Объединяем новую строку с текущим содержимым файла
    updatedContent := strings.Join(lines, "\n")

    // Устанавливаем позицию в начало файла
    file.Seek(0, 0)

    // Записываем обновленное содержимое в файл
    _, err = file.WriteString(updatedContent)
    if err != nil {
        fmt.Println(err)
        return nil
    }

	return nil
}

func rewriteRoutesConfig(filePath string, fileData *model.FileData) error {
	configName := fmt.Sprintf(`%vPageRoutesConfig`, fileData.PascalCaseFileName)
	importLine := fmt.Sprintf(`import { %v } from "../pages/%vPage";`, configName, fileData.PascalCaseFileName);

	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
    if err != nil {
        fmt.Println(err)
        return nil
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
    var lines []string
	lines = append(lines, fmt.Sprintf(`%v`, importLine) + "\n")

	isInsideRoutesArr := false;

    for scanner.Scan() {
		if strings.Contains(scanner.Text(), configName) {
			return fmt.Errorf("Module with this ROUTE config name already exist")
		}

		if strings.Contains(scanner.Text(), "const routeConfigs") {
			isInsideRoutesArr = true
		}

		if strings.Contains(scanner.Text(), "];") && isInsideRoutesArr {
			lines = append(lines, fmt.Sprintf("\t" + "%v,", configName))
			isInsideRoutesArr = false
		}

        lines = append(lines, scanner.Text())
    }

    // Объединяем новую строку с текущим содержимым файла
    updatedContent := strings.Join(lines, "\n")

    // Устанавливаем позицию в начало файла
    file.Seek(0, 0)

    // Записываем обновленное содержимое в файл
    _, err = file.WriteString(updatedContent)
    if err != nil {
        fmt.Println(err)
        return nil
    }

	return nil
}

func rewriteTranslateConfig(filePath string, fileData *model.FileData) error {
	configName := fmt.Sprintf(`%vMenuTranslate`, fileData.PascalCaseFileName)
	importLine := fmt.Sprintf(`import { %v } from "src/app/pages/%vPage";`, configName, fileData.PascalCaseFileName);

	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
    if err != nil {
        fmt.Println(err)
        return nil
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
    var lines []string
	lines = append(lines, fmt.Sprintf(`%v`, importLine) + "\n")

	isInsideRoutesArr := false;

    for scanner.Scan() {
		if strings.Contains(scanner.Text(), configName) {
			return fmt.Errorf("Module with this TRANSLATE config name already exist")
		}

		if strings.Contains(scanner.Text(), "const locale") {
			isInsideRoutesArr = true
		}

		if strings.Contains(scanner.Text(), "};") && isInsideRoutesArr {
			lines = append(lines, fmt.Sprintf("\t" + "...%v,", configName))
			isInsideRoutesArr = false
		}

        lines = append(lines, scanner.Text())
    }

    // Объединяем новую строку с текущим содержимым файла
    updatedContent := strings.Join(lines, "\n")

    // Устанавливаем позицию в начало файла
    file.Seek(0, 0)

    // Записываем обновленное содержимое в файл
    _, err = file.WriteString(updatedContent)
    if err != nil {
        fmt.Println(err)
        return nil
    }

	return nil
}

func createFileAndWriteData(filePath string, data string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}

	content := []byte(data)

	if _, err := file.Write(content); err != nil {
		log.Fatal(err)
	}

	file.Sync()
	defer file.Close()
}

func createDir(dirPath string) error {
	if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func generateFoldersInPage(filePath string,folders []string) {
	for _, dirName := range folders {
		dirPath := filepath.Join(filePath, dirName);

		createDir(dirPath)
	}
}

func generateConfigContent(configFilePath string, newFileData *model.FileData) {
	fileName :=  fmt.Sprintf(`%vRoutesConfig`, newFileData.PascalCaseFileName);

	rootFiles := [1]model.CustomFileProperty{
		{Name: fileName, Extensions: []string{"jsx"}, BuilderFunc: fileconfig.CreateConfigFile},
	}

	for _, fileProps := range rootFiles {
		fullFileName := model.AddExtensions(fileProps.Name, fileProps.Extensions...)

		filePath := filepath.Join(configFilePath, fullFileName)

		fileConent := fileProps.BuilderFunc(*newFileData)

		createFileAndWriteData(filePath, fileConent)
	}
}

func generateConstsContent(configFilePath string, newFileData *model.FileData) {
	fileName := "urlPaths";

	rootFiles := [1]model.CustomFileProperty{
		{Name: fileName, Extensions: []string{"jsx"}, BuilderFunc: consts.CreateUrlPaths},
	}

	for _, fileProps := range rootFiles {
		fullFileName := model.AddExtensions(fileProps.Name, fileProps.Extensions...)

		filePath := filepath.Join(configFilePath, fullFileName)

		fileConent := fileProps.BuilderFunc(*newFileData)

		createFileAndWriteData(filePath, fileConent)
	}
}

func generateFormSchemaContent(configFilePath string, newFileData *model.FileData) {
	fileName := fmt.Sprintf(`%vFormSchema`, newFileData.CamelCaseFileName);

	rootFiles := [1]model.CustomFileProperty{
		{Name: fileName, Extensions: []string{"js"}, BuilderFunc: models.CreateFormSchema},
	}

	for _, fileProps := range rootFiles {
		fullFileName := model.AddExtensions(fileProps.Name, fileProps.Extensions...)

		filePath := filepath.Join(configFilePath, fullFileName)

		fileConent := fileProps.BuilderFunc(*newFileData)

		createFileAndWriteData(filePath, fileConent)
	}
}