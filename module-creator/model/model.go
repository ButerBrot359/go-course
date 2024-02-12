package model

import (
	"fmt"
	"strings"
)

type FileData struct {
	RootDirectory      string
	FolderName         string
	PascalCaseFileName string
	SnakeCaseFileName  string
	UppercaseTitle string
	CamelCaseFileName  string
	HasTabs            bool
}

type FileFormatter interface {
	SetPascalCase()
	SetSnakeCase()
	SetCamelCase()
	SetUppercaseTitle()
	AddPostfix(postfix string) string
}

func (fileData *FileData) SetPascalCase() {
	newName := ""

	words := strings.Split(fileData.FolderName, "-")

	for _, word := range words {
		titledWord := strings.Title(strings.ToLower(word))
		newName += titledWord
	}

	fileData.PascalCaseFileName = newName
}

func (fileData *FileData) SetCamelCase() {
	newName := ""

	words := strings.Split(fileData.FolderName, "-")

	for i, word := range words {
		if i == 0 {
			newName += word
			continue
		}
		titledWord := strings.Title(strings.ToLower(word))
		newName += titledWord
	}

	fileData.CamelCaseFileName = newName
}

func (fileData *FileData) SetUppercaseTitle() {
	var uppercaseArr []string
	words := strings.Split(fileData.FolderName, "-")

	for _, word := range words {
		uppercaseArr = append(uppercaseArr, strings.ToUpper(word))
	}

	fileData.UppercaseTitle = strings.Join(uppercaseArr[:], "_")
}

func (fileData *FileData) SetSnakeCase() {
	newNameArr := make([]string, 0)

	words := strings.Split(fileData.FolderName, "-")

	for _, word := range words {
		lowercaseWord := strings.ToLower(word)
		newNameArr = append(newNameArr, lowercaseWord)
	}

	fileData.SnakeCaseFileName = strings.Join(newNameArr, "-")
	fileData.FolderName = strings.Join(newNameArr, "-")
}

func (fileData *FileData) AddPostfix(postfix string) string {
	return fileData.PascalCaseFileName + strings.Title(strings.ToLower(postfix))
}

func AddExtensions(fullFileName string, extensions ...string) string {
	for _, ext := range extensions {
		fullFileName += fmt.Sprintf(".%v", ext)
	}

	return fullFileName
}

type FileProperty struct {
	Postfix     string
	Extensions  []string
	Type        string
	BuilderFunc func(string, FileData) string
}

type CustomFileProperty struct {
	Name        string
	Extensions  []string
	BuilderFunc func(FileData) string
}

const DefaultFileType = "default"
const TabsFileType = "tabs"

type PathConstructor interface {
	SetRootDirectory()
}

func (fileData *FileData) SetRootDirectory(path string) {
	fileData.RootDirectory = path
}
