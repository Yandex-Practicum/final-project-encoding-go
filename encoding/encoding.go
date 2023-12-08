package encoding

import (
	"encoding/json"
	"fmt"
	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
	"os"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {

	var dockerCompose models.DockerCompose

	jsonData, err := os.ReadFile(j.FileInput)

	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s\n", err.Error())
		return err
	}

	err = json.Unmarshal(jsonData, &dockerCompose)
	if err != nil {
		fmt.Printf("ошибка при декодировании содержимого jsonData: %s\n", err.Error())
		return err
	}

	yamlData, err := yaml.Marshal(&dockerCompose)

	if err != nil {
		fmt.Printf("ошибка при кодировании dockerCompose в формат YAML: %s\n", err.Error())
		return err
	}

	err = os.WriteFile(j.FileOutput, yamlData, 0644)

	if err != nil {
		fmt.Printf("ошибка при записи содержимого yamlData: %s\n", err.Error())
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	var dockerCompose models.DockerCompose

	yamlData, err := os.ReadFile(y.FileInput)

	if err != nil {
		fmt.Printf("произошла ошибка при чтении файла: %s\n", err.Error())
		return err
	}

	err = yaml.Unmarshal(yamlData, &dockerCompose)

	if err != nil {
		fmt.Printf("произошла ошибка декодирования содержимого yamlData: %s\n", err.Error())
		return err
	}

	jsonData, err := json.Marshal(&dockerCompose)

	if err != nil {
		fmt.Printf("произошла ошибка кодирования в формат JSON: %s\n", err.Error())
		return err
	}

	err = os.WriteFile(y.FileOutput, jsonData, 0644)

	if err != nil {
		fmt.Printf("ошибка при записи содержимого jsonData: %s\n", err.Error())
		return err
	}

	return nil
}
