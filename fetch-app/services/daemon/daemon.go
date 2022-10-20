package daemon

import (
	"github.com/fahmyabdul/golibs"
	"github.com/fahmyabdul/golibs/daemon"
	"github.com/fahmyabdul/efishery-task/fetch-app/configs"
	"github.com/fahmyabdul/efishery-task/fetch-app/services/daemon/handlers"
)

type Daemon struct{}

// Start : Starting Daemon
func (p Daemon) Start() error {
	daemon := daemon.New()
	daemon.SleepTime = configs.Properties.Services.Daemon.Sleep
	daemon.Logger = golibs.Log
	daemon.WaitGroup = configs.Properties.Services.Daemon.WaitGroup

	daemon.AddHandler((&handlers.Handlers{}).Handle)

	daemon.StartDaemon()

	return nil
}
