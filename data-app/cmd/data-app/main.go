package main

import (
	"sync"

	"github.com/fahmyabdul/golibs"
	"github.com/fahmyabdul/self-growth/data-app/app"
	"github.com/fahmyabdul/self-growth/data-app/configs"
	"github.com/fahmyabdul/self-growth/data-app/internal/initialization"
	"github.com/fahmyabdul/self-growth/data-app/services"
	"github.com/fahmyabdul/self-growth/data-app/services/prometheus"
	"github.com/fahmyabdul/self-growth/data-app/services/restapi"
)

func main() {
	defer app.Properties.ClosingApp()

	err := app.Initialize("Data App")
	if err != nil {
	ReInitialize:
		app.Sleeping()
		err := app.Initialize("Data App")
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
