package main

import (
	"github.com/fahmyabdul/self-growth/fetch-app/app"
	"github.com/fahmyabdul/self-growth/fetch-app/services"
	"github.com/fahmyabdul/self-growth/fetch-app/services/kafka/consumer"
	"github.com/fahmyabdul/golibs"
)

func main() {
	defer app.Properties.ClosingApp()

	err := app.Initialize("Fetch App Consumer")
	if err != nil {
	ReInitialize:
		app.Sleeping()
		err := app.Initialize("Fetch App Consumer")
		if err != nil {
			goto ReInitialize
		}
	}

	err = services.Run(consumer.KafkaConsumer{})
	if err != nil {
		golibs.Log.Printf("| Services | Error: %s\n", err.Error())
	}
}
