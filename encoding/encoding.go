package encoding

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3" // импортируем пакет для работы с YAML
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
	bytesJson, err := os.ReadFile(j.FileInput)
	if err != nil {
		return fmt.Errorf("Не могу открыть файл: %s, %w", j.FileInput, err.Error())
	}
	// десериализуем JSON в DockerCompose
	if err = json.Unmarshal(bytesJson, &j.DockerCompose); err != nil {
		return fmt.Errorf("Ошибка десериализации файла: %s, %w", j.FileInput, err.Error())
	}
	// сериализация в Yaml
	bytesYaml, err := yaml.Marshal(j.DockerCompose)
	if err != nil {
		return fmt.Errorf("Ошибка сериализации Yaml: %w", err.Error())
	}
	//
	if err = os.WriteFile(j.FileOutput, bytesYaml, 0666); err != nil {
		return fmt.Errorf("yaml file '%s' creation fail: %w", j.FileOutput, err.Error())
	}
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	bytesYaml, err := os.ReadFile(y.FileInput)
	if err != nil {
		return fmt.Errorf("Не могу открыть файл: %s, %w", y.FileInput, err.Error())
	}
	// десериализуем Yaml в DockerCompose
	if err = yaml.Unmarshal(bytesYaml, &y.DockerCompose); err != nil {
		return fmt.Errorf("Ошибка десериализации файла '%s': %w", y.FileInput, err.Error())
	}
	// сериализация в Json
	bytesJson, err := json.Marshal(y.DockerCompose)
	if err != nil {
		return fmt.Errorf("Ошибка сериализации Yaml: %w", err.Error())
	}
	//
	if err = os.WriteFile(y.FileOutput, bytesJson, 0666); err != nil {
		return fmt.Errorf("json file '%s' creation fail: %w", y.FileOutput, err.Error())
	}
	return nil
}
