package encoding

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
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
	// ниже реализуйте метод
	jsonFile, err := os.ReadFile(j.FileInput)
	if err != nil {
		fmt.Println("Произошла ошибка чтения JSON-файла, err=", err)
		return err
	}

	err = json.Unmarshal(jsonFile, &j.DockerCompose)
	if err != nil {
		fmt.Println("Произошла ошибка десериализации JSON в структуру DockerCompose, err=", err)
		return err
	}

	yamlFile, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		fmt.Println("Произошла ошибка сериализации структуры DockerCompose в YAML, err=", err)
		return err
	}

	err = os.WriteFile(j.FileOutput, yamlFile, 0644)
	if err != nil {
		fmt.Println("Произошла ошибка записи YAML-файла, err=", err)
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод
	yamlFile, err := os.ReadFile(y.FileInput)
	if err != nil {
		fmt.Println("Произошла ошибка чтения YAML-файла, err=", err)
		return err
	}

	err = yaml.Unmarshal(yamlFile, &y.DockerCompose)
	if err != nil {
		fmt.Println("Произошла ошибка десериализации YAML в структуру DockerCompose, err=", err)
		return err
	}

	jsonFile, err := json.Marshal(&y.DockerCompose)
	if err != nil {
		fmt.Println("Произошла ошибка сериализации структуры DockerCompose в JSON, err=", err)
		return err
	}

	err = os.WriteFile(y.FileOutput, jsonFile, 0644)
	if err != nil {
		fmt.Println("Произошла ошибка записи JSON-файла, err=", err)
		return err
	}

	return nil
}
