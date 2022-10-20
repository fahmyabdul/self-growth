package main

import (
	"github.com/fahmyabdul/golibs"

	"github.com/fahmyabdul/efishery-task/fetch-app/app"
	"github.com/fahmyabdul/efishery-task/fetch-app/services"
	"github.com/fahmyabdul/efishery-task/fetch-app/services/restapi"
)

func main() {
	defer app.Properties.ClosingApp()

	err := app.Initialize("Fetch App RestAPI")
	if err != nil {
	ReInitialize:
		app.Sleeping()
		err := app.Initialize("Fetch App RestAPI")
		if err != nil {
			goto ReInitialize
		}
	}

	err = services.Run(restapi.Restapi{})
	if err != nil {
		golibs.Log.Printf("| Services | Error: %s\n", err.Error())
	}
}
