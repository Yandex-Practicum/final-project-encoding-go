package main

import (
	"fmt"

	"github.com/Yandex-Practicum/final-project-encoding-go/encoding"
	"github.com/Yandex-Practicum/final-project-encoding-go/utils"
)

func Encode(data encoding.MyEncoder) error {
	return data.Encoding()
}

func main() {
	if err := utils.CreateJSONFile(); err != nil {
		fmt.Printf("При генерации JSON-файла произошла ошибка: %s", err.Error())
		return
	}

	if err := utils.CreateYAMLFile(); err != nil {
		fmt.Printf("При генерации YAML-файла произошла ошибка: %s", err.Error())
		return
	}

	jsonData := encoding.JSONData{FileInput: "jsonInput.json", FileOutput: "yamlOutput.yml"}
	err := Encode(&jsonData)
	if err != nil {
		fmt.Printf("ошибка при перекодировании данных из JSON в YAML: %s", err.Error())
	}

	yamlData := encoding.YAMLData{FileInput: "yamlInput.yml", FileOutput: "jsonOutput.json"}
	err = Encode(&yamlData)
	if err != nil {
		fmt.Printf("ошибка при перекодировании данных из YAML в JSON: %s", err.Error())
	}
}
