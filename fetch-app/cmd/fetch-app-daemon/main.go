package main

import (
	"github.com/fahmyabdul/efishery-task/fetch-app/app"
	"github.com/fahmyabdul/efishery-task/fetch-app/services"
	"github.com/fahmyabdul/efishery-task/fetch-app/services/daemon"
	"github.com/fahmyabdul/golibs"
)

func main() {
	defer app.Properties.ClosingApp()

	err := app.Initialize("Fetch App Daemon")
	if err != nil {
	ReInitialize:
		app.Sleeping()
		err := app.Initialize("Fetch App Daemon")
		if err != nil {
			goto ReInitialize
		}
	}

	err = services.Run(daemon.Daemon{})
	if err != nil {
		golibs.Log.Printf("| Services | Error: %s\n", err.Error())
	}
}
