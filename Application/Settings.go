package application

import (
	"encoding/json"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func (ApplicationSettings *Settings) ConnectPostgreSQL() (DataBaseConnection *gorm.DB, Error error) {
	DataBaseConnection, Error = gorm.Open(postgres.Open(fmt.Sprintf("postgres://%s:%s@%s:%d", ApplicationSettings.PostgreSQLLogin, ApplicationSettings.PostgreSQLPassword, ApplicationSettings.PostgreSQLAdress, ApplicationSettings.PostgreSQLPort)), &gorm.Config{})
	if Error != nil {
		return DataBaseConnection, Error
	}
	return DataBaseConnection, Error
}
