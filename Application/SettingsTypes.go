package application

import "gorm.io/gorm"

type Settings struct {
	ServerAdress         string
	ServerPort           int32
	PostgreSQLAdress     string
	PostgreSQLPort       int32
	PostgreSQLLogin      string
	PostgreSQLPassword   string
	PostgreSQLConnection *gorm.DB
}
