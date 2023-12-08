package encoding

import (
	"encoding/json"
	"fmt"
	"io"
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

	jsonFile, err := os.Open(j.FileInput)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	if err = json.Unmarshal(jsonData, &j.DockerCompose); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	yamlData, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	yamlFile, err := os.Create(j.FileOutput)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer yamlFile.Close()

	_, err = yamlFile.Write(yamlData)
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {

	yamlFile, err := os.Open(y.FileInput)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer yamlFile.Close()
	yamlData, err := io.ReadAll(yamlFile)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	if err = yaml.Unmarshal(yamlData, &y.DockerCompose); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	jsonData, err := json.Marshal(&y.DockerCompose)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	jsonFile, err := os.Create(y.FileOutput)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}
