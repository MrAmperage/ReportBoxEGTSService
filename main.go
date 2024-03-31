package main

import (
	"fmt"

	Application "github.com/MrAmperage/ReportBoxEGTSService/Application"
)

func main() {
	Application := &Application.Application{}
	Error := Application.Start()
	if Error != nil {
		fmt.Println(Error)
	}

}
