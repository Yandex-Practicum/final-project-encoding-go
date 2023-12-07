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

	jsonFile, err := os.ReadFile(j.FileInput)
	if err != nil {
		return fmt.Errorf("ошибка при чтении %s файла: %w\n", j.FileInput, err)
	}

	err = json.Unmarshal(jsonFile, &j.DockerCompose)
	if err != nil {
		return fmt.Errorf("ошибка при десериализации %s файла: %w\n", j.FileInput, err)
	}
	yamlFile, err := os.Create(j.FileOutput)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %s: %w\n", j.FileOutput, err)
	}

	defer yamlFile.Close()

	yamlData, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		return fmt.Errorf("ошибка при сериализации в yaml: %w\n", err)
	}

	_, err = yamlFile.Write(yamlData)
	if err != nil {
		return fmt.Errorf("ошибка при записи данных в файл: %s: %w\n", j.FileOutput, err)
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	yamlFile, err := os.ReadFile(y.FileInput)
	if err != nil {
		return fmt.Errorf("ошибка при чтении %s файла: %w\n", y.FileInput, err)
	}

	err = yaml.Unmarshal(yamlFile, &y.DockerCompose)
	if err != nil {
		return fmt.Errorf("ошибка при десериализации %s файла: %w\n", y.FileInput, err)
	}

	jsonFile, err := os.Create(y.FileOutput)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %s: %w\n", y.FileOutput, err)
	}

	defer jsonFile.Close()

	jsonData, err := json.Marshal(&y.DockerCompose)
	if err != nil {
		return fmt.Errorf("ошибка при сериализации в json: %w\n", err)
	}

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		return fmt.Errorf("ошибка при записи данных в файл: %s: %w\n", y.FileOutput, err)
	}

	return nil
}
