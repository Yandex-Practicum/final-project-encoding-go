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
	// ниже реализуйте метод
	jsonFile, err := os.ReadFile(j.FileInput)
	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s", err.Error())
	}

	err = json.Unmarshal(jsonFile, &j.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при десериализации из json: %s\n", err.Error())
	}

	out, err := yaml.Marshal(j.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при сериализации в yaml: %s\n", err.Error())
	}

	f, err := os.Create("yamlOutput.yaml")
	if err != nil {
		fmt.Printf("ошибка при создании файла: %s", err.Error())
	}
	defer f.Close()

	_, err = f.Write(out)
	if err != nil {
		fmt.Printf("ошибка при записи данных в файл: %s", err.Error())
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод

	yamlFile, err := os.ReadFile(y.FileInput)
	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s", err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &y.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при десериализации из yaml: %s\n", err.Error())
	}

	out, err := json.Marshal(y.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при сериализации в json: %s\n", err.Error())
	}

	f, err := os.Create("jsonOutput.json")
	if err != nil {
		fmt.Printf("ошибка при создании файла: %s", err.Error())
	}
	defer f.Close()

	_, err = f.Write(out)
	if err != nil {
		fmt.Printf("ошибка при записи данных в файл: %s", err.Error())
	}

	return nil
}
