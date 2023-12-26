package encoding

import (
	"encoding/json"
	"fmt"
	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
	"io"
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
	jsonFile, err := os.Open(j.FileInput)
	if err != nil {
		fmt.Println("Не удалось открыть файл")
		return err
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println("Не удалось закрыть файл")
		}
	}(jsonFile)
	byteValue, _ := io.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &j.DockerCompose)
	if err != nil {
		fmt.Println("Не удалось провести десериализацию")
		return err
	}
	file, err := os.Create(j.FileOutput)
	if err != nil {
		fmt.Println("Не удалось создать файл")
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Не удалось закрыть файл")
		}
	}(file)
	yamlData, err := yaml.Marshal(j.DockerCompose)
	if err != nil {
		fmt.Println("Не удалось провести сериализацию")
		return err
	}
	_, err = file.Write(yamlData)
	if err != nil {
		fmt.Println("Не удалось записать данные в файл")
		return err
	}
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод
	// ...
	yamlFile, err := os.Open(y.FileInput)
	if err != nil {
		fmt.Println("Не удалось открыть файл")
		return err
	}
	defer func(yamlFile *os.File) {
		err := yamlFile.Close()
		if err != nil {
			fmt.Println("Не удалось закрыть файл")
		}
	}(yamlFile)
	byteValue, _ := io.ReadAll(yamlFile)
	err = yaml.Unmarshal(byteValue, &y.DockerCompose)
	if err != nil {
		fmt.Println("Не удалось провести десериализацию")
		return err
	}
	file, err := os.Create(y.FileOutput)
	if err != nil {
		fmt.Println("Не удалось создать файл")
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Не удалось закрыть файл")
		}
	}(file)
	jsonData, err := json.Marshal(y.DockerCompose)
	if err != nil {
		fmt.Println("Не удалось провести сериализацию")
		return err
	}
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Не удалось записать данные в файл")
		return err
	}
	return nil
}
