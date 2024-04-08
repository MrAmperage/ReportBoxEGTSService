package application

import (
	"fmt"
	"net"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*Запуск приложения*/
func (Application *Application) Start() (Error error) {
	Application.Settings, Error = Application.Settings.ReadSettings("Config.json")
	if Error != nil {
		return Error
	}
	Application.Settings.PostgreSQLConnection, Error = Application.ConnectPostgreSQL()
	if Error != nil {
		return Error
	}
	Error = Application.StartServer(func(Connection net.Conn) {
		fmt.Println("Пришел пакет")
	})
	if Error != nil {
		return Error
	}

	return Error
}
func (Application *Application) StartServer(Handler func(Connection net.Conn)) (Error error) {

	Application.Server, Error = net.Listen(Application.Settings.ServerType, fmt.Sprintf("%s:%d", Application.Settings.ServerAdress, Application.Settings.ServerPort))
	if Error != nil {
		return Error
	}

	for {
		Connection, Error := Application.Server.Accept()
		if Error != nil {
			return Error

		}
		go Handler(Connection)
	}

}

func (Application *Application) ConnectPostgreSQL() (DataBaseConnection *gorm.DB, Error error) {
	DataBaseConnection, Error = gorm.Open(postgres.Open(fmt.Sprintf("postgres://%s:%s@%s:%d", Application.Settings.PostgreSQLLogin, Application.Settings.PostgreSQLPassword, Application.Settings.PostgreSQLAdress, Application.Settings.PostgreSQLPort)), &gorm.Config{})
	if Error != nil {
		return DataBaseConnection, Error
	}
	return DataBaseConnection, Error
}
