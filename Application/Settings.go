package application

import (
	"encoding/json"
	"os"
)

/*Считывание настроек из файла*/
func (Settings *Settings) ReadSettings(FileName string) (Error error) {
	FileByte, Error := os.ReadFile(FileName)
	if Error != nil {
		return Error
	}
	Error = json.Unmarshal(FileByte, &Settings)
	if Error != nil {
		return Error
	}
	return Error

}
