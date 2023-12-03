package encoding

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"os"

	"github.com/AlexeyMurenkov/final-project-encoding-go/models"
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
	bytes, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &j.DockerCompose)
	if err != nil {
		return err
	}

	enc, err := yaml.Marshal(j.DockerCompose)
	if err != nil {
		return err
	}

	yamlFile, err := os.Create(j.FileOutput)
	_, err = yamlFile.Write(enc)
	if err != nil {
		return err
	}
	return yamlFile.Close()
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	bytes, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(bytes, &y.DockerCompose)
	if err != nil {
		return err
	}

	enc, err := json.Marshal(y.DockerCompose)
	if err != nil {
		return err
	}

	yamlFile, err := os.Create(y.FileOutput)
	_, err = yamlFile.Write(enc)
	if err != nil {
		return err
	}
	return yamlFile.Close()
}
