package encoding

import (
	"encoding/json"
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
	dataFile, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}

	var dockerCompose models.DockerCompose
	err = json.Unmarshal(dataFile, &dockerCompose)
	if err != nil {
		return err
	}

	dataFile, err = yaml.Marshal(dockerCompose)
	if err != nil {
		return err
	}

	err = os.WriteFile(j.FileOutput, dataFile, 0664)
	if err != nil {
		return err
	}
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод
	dataFile, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}

	var dockerCompose models.DockerCompose
	err = yaml.Unmarshal(dataFile, &dockerCompose)
	if err != nil {
		return err
	}

	dataFile, err = json.Marshal(dockerCompose)
	if err != nil {
		return err
	}

	err = os.WriteFile(y.FileOutput, dataFile, 0664)
	if err != nil {
		return err
	}
	return nil
}
