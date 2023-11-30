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
	var dockerCompose models.DockerCompose

	jsonFile, err := os.Open(j.FileInput)

	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	if err = json.Unmarshal(byteValue, &dockerCompose); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	yamlData, err := yaml.Marshal(&dockerCompose)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	yamlFile, err := os.Create(j.FileOutput)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}

	_, err = yamlFile.Write(yamlData)
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	var dockerCompose models.DockerCompose

	yamlFile, err := os.Open(y.FileInput)

	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}

	defer yamlFile.Close()

	byteValue, err := io.ReadAll(yamlFile)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	if err = yaml.Unmarshal(byteValue, &dockerCompose); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	jsonData, err := json.Marshal(&dockerCompose)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	jsonFile, err := os.Create(y.FileOutput)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}
