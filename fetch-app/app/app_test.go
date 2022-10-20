package app

import (
	"log"
	"testing"

	"github.com/fahmyabdul/efishery-task/fetch-app/configs"
	"github.com/fahmyabdul/golibs"
)

func TestDatabaseConnection(t *testing.T) {
	golibs.Log = log.Default()

	structApp := Application{}

	err := configs.InitConfig("../.config.example.yaml", "TESTING")
	if err != nil {
		t.Errorf("Failed InitConfig, err : %s\n", err.Error())
	}

	err = structApp.DatabaseConnection()
	if err != nil {
		t.Errorf("Failed DatabaseConnection, err : %s\n", err.Error())
	}
}
