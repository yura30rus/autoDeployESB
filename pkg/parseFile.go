package pkg

import (
	"log"
	"os"
	"strings"
)

type App struct {
	Command []string
	AppName string
}

func ReadFile(pathToFile string) []App {
	content, err := os.ReadFile(pathToFile)
	result := make([]App, 0)
	if err != nil {
		log.Fatal(err)
	}
	allCommandsFromFile := strings.Split(string(content), "===")
	for i := range allCommandsFromFile {
		oneCommand := strings.Split(allCommandsFromFile[i], "\n")
		oneApp := App{make([]string, 0), ""}
		for _, elem := range oneCommand {
			if elem == "" || elem == "\r" {
				continue
			}
			if strings.Contains(elem, "app name :") || elem == "" || elem == "\r" {
				oneApp.AppName = strings.Replace(elem, "app name : ", "", -1)
				continue
			}
			oneApp.Command = append(oneApp.Command, elem)
		}
		result = append(result, oneApp)
	}
	return result
}
