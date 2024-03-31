package application

/*Запуск приложения*/
func (Application *Application) Start() (Error error) {
	Application.Settings, Error = Application.Settings.ReadSettings("Config.json")
	if Error != nil {
		return Error
	}
	Application.Settings.PostgreSQLConnection, Error = Application.Settings.ConnectPostgreSQL()
	if Error != nil {
		return Error
	}
	return Error
}
