package main

import (
	"sync"

	"github.com/fahmyabdul/self-growth/fetch-app/app"
	"github.com/fahmyabdul/self-growth/fetch-app/configs"
	"github.com/fahmyabdul/self-growth/fetch-app/internal/initialization"
	"github.com/fahmyabdul/self-growth/fetch-app/services"
	"github.com/fahmyabdul/self-growth/fetch-app/services/cronjob"
	"github.com/fahmyabdul/self-growth/fetch-app/services/daemon"
	"github.com/fahmyabdul/self-growth/fetch-app/services/kafka/consumer"
	"github.com/fahmyabdul/self-growth/fetch-app/services/kafka/publisher"
	"github.com/fahmyabdul/self-growth/fetch-app/services/prometheus"
	"github.com/fahmyabdul/self-growth/fetch-app/services/restapi"
	"github.com/fahmyabdul/golibs"
)

// @Security ApiKeyAuth
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
func main() {
	defer app.Properties.ClosingApp()

	err := app.Initialize("Fetch App")
	if err != nil {
	ReInitialize:
		app.Sleeping()
		err := app.Initialize("Fetch App")
		if err != nil {
			goto ReInitialize
		}
	}

ReInitBefore:
	// Initialize for contents before Services started
	err = initialization.InitBefore()
	if err != nil {
		app.Sleeping()
		golibs.Log.Printf("| InitBefore | Error: %s\n", err.Error())
		goto ReInitBefore
	}

	wg := &sync.WaitGroup{}

	if configs.Properties.Services.Prometheus.Status {
		wg.Add(1)
		go func() {
			err = services.Run(prometheus.Prometheus{})
			if err != nil {
				golibs.Log.Printf("| Services | Error: %s\n", err.Error())
			}
			wg.Done()
		}()
	}

	if configs.Properties.Services.Kafka.Publisher.Status {
		wg.Add(1)
		err = services.Run(publisher.KafkaPublisher{})
		if err != nil {
			golibs.Log.Printf("| Services | Error: %s\n", err.Error())
		}
		wg.Done()
	}

	if configs.Properties.Services.Kafka.Consumer.Status {
		wg.Add(1)
		go func() {
			err = services.Run(consumer.KafkaConsumer{})
			if err != nil {
				golibs.Log.Printf("| Services | Error: %s\n", err.Error())
			}
			wg.Done()
		}()
	}

	if configs.Properties.Services.CronJob.Status {
		wg.Add(1)
		go func() {
			err = services.Run(cronjob.CronJob{})
			if err != nil {
				golibs.Log.Printf("| Services | Error: %s\n", err.Error())
			}
			wg.Done()
		}()
	}

	if configs.Properties.Services.Daemon.Status {
		wg.Add(1)
		go func() {
			err = services.Run(daemon.Daemon{})
			if err != nil {
				golibs.Log.Printf("| Services | Error: %s\n", err.Error())
			}
			wg.Done()
		}()
	}

	if configs.Properties.Services.Restapi.Status {
		wg.Add(1)
		go func() {
			err = services.Run(restapi.Restapi{})
			if err != nil {
				golibs.Log.Printf("| Services | Error: %s\n", err.Error())
			}
			wg.Done()
		}()
	}

ReInitAfter:
	// Initialize for contents that requires Services to be started first
	err = initialization.InitAfter()
	if err != nil {
		golibs.Log.Printf("| InitAfter | Error: %s\n", err.Error())
		goto ReInitAfter
	}

	// Set Gauge to 1 means app successfully started
	prometheus.AppStatusGauge.Set(1)
	golibs.Log.Println("....Successfully starting Fetch App....")

	wg.Wait()
}
