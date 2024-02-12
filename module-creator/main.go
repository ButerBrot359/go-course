package main

import (
	"bufio"
	"consts"
	"fileconfig"
	"flag"
	"fmt"
	"log"
	"modals"
	"model"
	"models"
	"os"
	"path/filepath"
	"strings"
	"templates"
	"ui"
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

	pageModuleName := fmt.Sprintf(`%vPage`, newFileData.PascalCaseFileName);
	moduleDirPath := filepath.Join(newFileData.RootDirectory,"src", "app", "pages", pageModuleName)

	configDirPath := filepath.Join(newFileData.RootDirectory, "src", "app", "configs")

	if err := rewriteConfigFiles(configDirPath, newFileData); err != nil {
		fmt.Println(err)
		return
	}

	if err := createDir(moduleDirPath); err != nil {
		log.Fatal("Folder already exist")
		return
	}

	folders := []string{"api", "config", "consts", "models", "ui", "utils"}
	generateFoldersInPage(moduleDirPath, folders)

	pageUiFolder := filepath.Join(moduleDirPath, "ui", fmt.Sprintf(`%vPage`, newFileData.PascalCaseFileName))
	if err := createDir(pageUiFolder); err != nil {
		log.Fatal("Folder already exist")
		return
	}

	pageModalsFolder := filepath.Join(moduleDirPath, "ui", "Modals")
	if err := createDir(pageModalsFolder); err != nil {
		log.Fatal("Folder already exist")
		return
	}

	pageModalForm := filepath.Join(pageModalsFolder, fmt.Sprintf(`%vForm`, newFileData.PascalCaseFileName))
	if err := createDir(pageModalForm); err != nil {
		log.Fatal("Folder already exist")
		return
	}

	pageModalView := filepath.Join(pageModalsFolder, fmt.Sprintf(`%vView`, newFileData.PascalCaseFileName))
	if err := createDir(pageModalView); err != nil {
		log.Fatal("Folder already exist")
		return
	}

	generateConfigContent(filepath.Join(moduleDirPath, "config"), newFileData)
	generateConstsContent(filepath.Join(moduleDirPath, "consts"), newFileData)
	generateFormSchemaContent(filepath.Join(moduleDirPath, "models"), newFileData)
	generateUiPage(pageUiFolder, newFileData)
	generateIndexFile(moduleDirPath, newFileData)
	generateModalForm(pageModalForm, newFileData)
	generateModalView(pageModalView, newFileData)

}


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

func generateIndexFile(configFilePath string, newFileData *model.FileData) {
	fileName := "index"

	rootFiles := [1]model.CustomFileProperty{
		{Name: fileName, Extensions: []string{"js"}, BuilderFunc: templates.CreateIndexFile},
	}

	for _, fileProps := range rootFiles {
		fullFileName := model.AddExtensions(fileProps.Name, fileProps.Extensions...)

		filePath := filepath.Join(configFilePath, fullFileName)

		fileConent := fileProps.BuilderFunc(*newFileData)

		createFileAndWriteData(filePath, fileConent)
	}
}

func generateUiPage(configFilePath string, newFileData *model.FileData) {
	pageFileName := fmt.Sprintf(`%vPage`, newFileData.PascalCaseFileName);
	contentFileName := fmt.Sprintf(`%vContent`, newFileData.PascalCaseFileName);


	rootFiles := [3]model.CustomFileProperty{
		{Name: pageFileName, Extensions: []string{"jsx"}, BuilderFunc: ui.CreatePage},
		{Name: contentFileName, Extensions: []string{"jsx"}, BuilderFunc: ui.CreatePageContent},
		{Name: contentFileName, Extensions: []string{"module", "scss"}, BuilderFunc: ui.CreateContentStyles},

	}

	for _, fileProps := range rootFiles {
		fullFileName := model.AddExtensions(fileProps.Name, fileProps.Extensions...)

		filePath := filepath.Join(configFilePath, fullFileName)

		fileConent := fileProps.BuilderFunc(*newFileData)

		createFileAndWriteData(filePath, fileConent)
	}
}

func generateModalForm(configFilePath string, newFileData *model.FileData) {
	pageFileName := fmt.Sprintf(`%vForm`, newFileData.PascalCaseFileName);

	rootFiles := [2]model.CustomFileProperty{
		{Name: pageFileName, Extensions: []string{"jsx"}, BuilderFunc: modals.CreateModalForm},
		{Name: pageFileName, Extensions: []string{"module", "scss"}, BuilderFunc: modals.CreateModalFormStyle},
	}

	for _, fileProps := range rootFiles {
		fullFileName := model.AddExtensions(fileProps.Name, fileProps.Extensions...)

		filePath := filepath.Join(configFilePath, fullFileName)

		fileConent := fileProps.BuilderFunc(*newFileData)

		createFileAndWriteData(filePath, fileConent)
	}
}

func generateModalView(configFilePath string, newFileData *model.FileData) {
	pageFileName := fmt.Sprintf(`%vView`, newFileData.PascalCaseFileName);

	rootFiles := [2]model.CustomFileProperty{
		{Name: pageFileName, Extensions: []string{"jsx"}, BuilderFunc: modals.CreateModalView},
		{Name: pageFileName, Extensions: []string{"module", "scss"}, BuilderFunc: modals.CreateModalViewStyle},
	}

	for _, fileProps := range rootFiles {
		fullFileName := model.AddExtensions(fileProps.Name, fileProps.Extensions...)

		filePath := filepath.Join(configFilePath, fullFileName)

		fileConent := fileProps.BuilderFunc(*newFileData)

		createFileAndWriteData(filePath, fileConent)
	}
}