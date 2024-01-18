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
	// ниже реализуйте метод
	// ...

	// Читаем данные из файла JSON
	jsonFile, err := os.ReadFile(j.FileInput)
	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s", err.Error())
		return err
	}

	err = json.Unmarshal(jsonFile, &j.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при десериализации: %s", err.Error())
		return err
	}

	// Сохраняем в файл YAML
	yamlFile, err := os.Create(j.FileOutput)
	if err != nil {
		fmt.Printf("json file creation fail: %s", err.Error())
	}

	defer yamlFile.Close() // когда программа завершится, надо закрыть дескриптор файла

	out, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		fmt.Printf("json encoding fail: %s", err.Error())
	}

	_, err = yamlFile.Write(out)
	if err != nil {
		fmt.Printf("writing data fail: %s", err.Error())
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод
	// ...

	// Читаем данные из файла YAML
	yamlFile, err := os.ReadFile(y.FileInput)
	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s", err.Error())
		return err
	}

	err = yaml.Unmarshal(yamlFile, &y.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при десериализации: %s", err.Error())
		return err
	}

	// Сохраняем в файл JSON
	jsonFile, err := os.Create(y.FileOutput)
	if err != nil {
		fmt.Printf("json file creation fail: %s", err.Error())
	}

	defer jsonFile.Close() // когда программа завершится, надо закрыть дескриптор файла

	out, err := json.MarshalIndent(&y.DockerCompose, "", "    ")
	if err != nil {
		fmt.Printf("json encoding fail: %s", err.Error())
	}

	_, err = jsonFile.Write(out)
	if err != nil {
		fmt.Printf("writing data fail: %s", err.Error())
	}

	return nil
}
