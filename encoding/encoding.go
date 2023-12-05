package encoding

import (
	"encoding/json"
	"errors"
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
	//
	/* Комментарий студента:
	* Err ладно, она не нужна дальше своего блока
	* а вот как не выносить Data, если в блоке ReadFile обьявляется
	* то в блоке Unmarshal попросту не видна
	* отсюда и начал выносить
	* Получается надо вносить блок в блок или делать else,
	* что кажется намного хуже чем выносить переменную
	if bytesJson, err := os.ReadFile(j.FileInput); err != nil {
		return errors.New(fmt.Sprintf("Не могу открыть файл: %s, %s", j.FileInput, err.Error()))
	}
	// десериализуем JSON в DockerCompose
	if err := json.Unmarshal(bytesJson, j.DockerCompose); err != nil {
		return errors.New("Ошибка десериализации файла: " + j.FileInput)
	}
	*/
	var data []byte
	var err error
	//
	if data, err = os.ReadFile(j.FileInput); err != nil {
		return errors.New(fmt.Sprintf("Не могу открыть файл: %s, %s", j.FileInput, err.Error()))
	}
	// десериализуем JSON в DockerCompose
	if err := json.Unmarshal(data, &j.DockerCompose); err != nil {
		return errors.New("Ошибка десериализации файла: " + j.FileInput)
	}
	// сериализация в Yaml
	if data, err = yaml.Marshal(j.DockerCompose); err != nil {
		return errors.New(fmt.Sprintf("Ошибка сериализации Yaml: %s", err.Error()))
	}
	//
	if err = os.WriteFile(j.FileOutput, data, 0666); err != nil {
		return errors.New(fmt.Sprintf("yaml file '%s' creation fail: %s", j.FileOutput, err.Error()))
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	var data []byte
	var err error
	//
	if data, err = os.ReadFile(y.FileInput); err != nil {
		return errors.New(fmt.Sprintf("Не могу открыть файл: %s, %s", y.FileInput, err.Error()))
	}
	// десериализуем Yaml в DockerCompose
	if err = yaml.Unmarshal(data, &y.DockerCompose); err != nil {
		return errors.New("Ошибка десериализации файла: " + y.FileInput)
	}
	// сериализация в Json
	if data, err = json.Marshal(y.DockerCompose); err != nil {
		return errors.New(fmt.Sprintf("Ошибка сериализации Yaml: %s", err.Error()))
	}
	//
	if err = os.WriteFile(y.FileOutput, data, 0666); err != nil {
		return errors.New(fmt.Sprintf("json file '%s' creation fail: %s", y.FileOutput, err.Error()))
	}

	return nil
}
