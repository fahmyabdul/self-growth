package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/fahmyabdul/golibs"
	"github.com/fahmyabdul/self-growth/golang-architecture/configs"
)

type Application struct {
	Databases Databases
}

const (
	CurrentVersion string = "v2022.11.10-1"
)

var (
	Properties  Application
	ServiceName string
)

func GetVersion() string {
	fmt.Println(CurrentVersion)
	return CurrentVersion
}

// Initialize : args[0] -> Service Name
func Initialize(args ...string) error {
	if getVersion != false {
		GetVersion()
		os.Exit(0)
	}

	if len(args) > 0 {
		ServiceName = args[0]
	}

	err := configs.InitConfig(ConfigPath, ServiceName)
	if err != nil {
		log.Println("| InitConfig |", err.Error())
		return err
	}

	err = SetLog(ServiceName, LogPath)
	if err != nil {
		log.Println("| SetLog |", err.Error())
		return err
	}

	if ServiceName != "" {
		golibs.Log.Printf("....Starting %s %s....\n", args[0], CurrentVersion)
	}

	err = Properties.DatabaseConnection()
	if err != nil {
		return err
	}

	CloseHandler()

	return nil
}

func SetLog(serviceName, logPath string) error {
	if logPath == "" {
		logPath = fmt.Sprintf("%s%s%s", "./log/", strings.ToLower(strings.ReplaceAll(serviceName, " ", "_")), ".log")
	}

	err := (&golibs.LumberjackLogger{
		LogPath:       logPath,
		DailyRotate:   configs.Properties.Logger.DailyRotate,
		CompressLog:   configs.Properties.Logger.CompressLog,
		LogToTerminal: configs.Properties.Logger.LogToTerminal,
	}).SetLog()
	if err != nil {
		return err
	}

	return nil
}

// ClosingApp :
func (a *Application) ClosingApp() {
	a.DatabaseConnectionClose()
}

// CloseHandler :
func CloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		Properties.ClosingApp()
		os.Exit(0)
	}()
}
