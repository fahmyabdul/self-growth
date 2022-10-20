package version

import (
	"log"
	"testing"

	"github.com/fahmyabdul/golibs"
	"github.com/fahmyabdul/efishery-task/fetch-app/app"
)

func TestGetVersion(t *testing.T) {
	golibs.Log = log.Default()

	resp := GetVersion()
	if resp.Version != app.CurrentVersion {
		t.Errorf("Version [%s] not match app.CurrentVersion [%s]\n", resp.Version, app.CurrentVersion)
	}

	if resp.Version == "" {
		t.Errorf("Version is empty\n")
	}
}
