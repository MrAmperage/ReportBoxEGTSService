package application

/*Запуск приложения*/
func (Application *Application) Start() {
	Application.Settings.ReadSettings("Config.json")
}
