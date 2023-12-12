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
	var f *os.File
	// прочитать содержимое файла JSON
	data, err := os.ReadFile(j.FileInput)
	if err != nil {
		fmt.Println("Ошибка чтения файла JSON: ", err)
		return err
	}
	// перекодировать содержимое файла в структуру данных
	err = json.Unmarshal(data, &j.DockerCompose)
	if err != nil {
		fmt.Println("Ошибка перекодирования JSON в данные: ", err)
		return err
	}
	// перекодировать структуру данных в YAML
	data, err = yaml.Marshal(&j.DockerCompose)
	if err != nil {
		fmt.Println("Ошибка перекодирования данных в YAML: ", err)
		return err
	}
	// проверить наличие файла, если нет создать
	if _, err = os.Stat(j.FileOutput); os.IsNotExist(err) {
		// создать
		f, err = os.Create(j.FileOutput)
		if err != nil {
			fmt.Println("Ошибка создания файла YAML: ", err)
			return err
		}
	} else {
		if f, err = os.OpenFile(j.FileOutput, os.O_APPEND|os.O_WRONLY, 0644); err != nil {
			fmt.Println("Ошибка открытия файла YAML: ", err)
			return err
		}
	}
	defer f.Close()
	// записать содержимое в файл
	_, err = f.Write(data)
	if err != nil {
		fmt.Println("Ошибка записи в файл YAML: ", err)
		return err
	}
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод
	// ниже реализуйте метод
	var f *os.File
	// прочитать содержимое файла YAML
	data, err := os.ReadFile(y.FileInput)
	if err != nil {
		fmt.Println("Ошибка чтения файла YAML: ", err)
		return err
	}
	// перекодировать содержимое файла в структуру данных
	err = yaml.Unmarshal(data, &y.DockerCompose)
	if err != nil {
		fmt.Println("Ошибка перекодирования YAML в данные: ", err)
		return err
	}
	// перекодировать структуру данных в YAML
	data, err = json.Marshal(&y.DockerCompose)
	if err != nil {
		fmt.Println("Ошибка перекодирования данных в JSON: ", err)
		return err
	}
	// проверить наличие файла, если нет создать
	if _, err = os.Stat(y.FileOutput); os.IsNotExist(err) {
		// создать
		f, err = os.Create(y.FileOutput)
		if err != nil {
			fmt.Println("Ошибка создания файла JSON: ", err)
			return err
		}
	} else {
		if f, err = os.OpenFile(y.FileOutput, os.O_APPEND|os.O_WRONLY, 0644); err != nil {
			fmt.Println("Ошибка открытия файла JSON: ", err)
			return err
		}
	}
	defer f.Close()
	// записать содержимое в файл
	_, err = f.Write(data)
	if err != nil {
		fmt.Println("Ошибка записи в файл JSON: ", err)
		return err
	}
	return nil
}
