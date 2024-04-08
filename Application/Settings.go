package application

import (
	"encoding/json"
	"os"
)

/*Считывание настроек из файла*/
func (ApplicationSettings *Settings) ReadSettings(FileName string) (Settings Settings, Error error) {
	FileByte, Error := os.ReadFile(FileName)
	if Error != nil {
		return Settings, Error
	}
	Error = json.Unmarshal(FileByte, &Settings)
	if Error != nil {
		return Settings, Error
	}
	return Settings, Error

}
