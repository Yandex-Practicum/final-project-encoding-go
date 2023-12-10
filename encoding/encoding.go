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
	// Чтение JSON файла.
	jsonData, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}

	// Десериализация JSON в объект
	var dockerCompose models.DockerCompose
	if err = json.Unmarshal(jsonData, &dockerCompose); err != nil {
		return err
	}

	// Сериализация в YAML
	yamlData, err := yaml.Marshal(dockerCompose)
	if err != nil {
		return err
	}

	// Создание YAML-файла
	f, err := os.Create(j.FileOutput)
	if err != nil {
		return err
	}
	defer f.Close()

	// Запись слайса байт в файл
	_, err = f.Write(yamlData)
	if err != nil {
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Чтение YAML файла.
	yamlData, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}

	// Десериализация YAML в объект
	var dockerCompose models.DockerCompose
	if err = yaml.Unmarshal(yamlData, &dockerCompose); err != nil {
		return err
	}

	// Сериализация в JSON
	jsonData, err := json.MarshalIndent(dockerCompose, "", "    ")
	if err != nil {
		return err
	}

	// Создание yaml-файла
	f, err := os.Create(y.FileOutput)
	if err != nil {
		return err
	}
	defer f.Close()

	// Запись слайса байт в файл
	_, err = f.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}
