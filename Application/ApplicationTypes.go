package application

import (
	"net"
)

type Application struct {
	Settings Settings
	Server   net.Listener
}
