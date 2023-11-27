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
	j.DockerCompose = new(models.DockerCompose)
	bytes, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, j.DockerCompose)
	if err != nil {
		return err
	}
	bytes, err = yaml.Marshal(j.DockerCompose)
	if err != nil {
		return err
	}
	os.WriteFile(j.FileOutput, bytes, 0666)
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	y.DockerCompose = new(models.DockerCompose)
	bytes, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bytes, y.DockerCompose)
	if err != nil {
		return err
	}
	bytes, err = json.Marshal(y.DockerCompose)
	if err != nil {
		return err
	}
	os.WriteFile(y.FileOutput, bytes, 0666)
	return nil
	return nil
}
