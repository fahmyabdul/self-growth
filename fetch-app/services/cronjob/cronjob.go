package cronjob

import (
	"github.com/fahmyabdul/self-growth/fetch-app/configs"
	"github.com/fahmyabdul/self-growth/fetch-app/internal/common"
	"github.com/fahmyabdul/golibs"
	"github.com/jasonlvhit/gocron"
)

type CronJob struct{}

type jobFunction func(interface{}) error

func (p CronJob) Start() error {
	config := configs.Properties.Services.CronJob
	// Loop every jobs from configs file
	for jobName, jobConfig := range config.Jobs {
		// Handler for cronJob based on job name
		switch jobName {
		case "exchangerates":
			// currentJob is the job(function) that will be executed by cronjob
			// the function needs to be in [ func(interface{}) error ] format
			currentJob := common.CheckExchangeRates
			p.DoJobs(jobName, jobConfig.Every, jobConfig.Hours, currentJob)
		default:
			golibs.Log.Printf("| CronJob | Job: %s, is unrecognized\n", jobName)
			continue
		}
	}

	<-gocron.Start()

	return nil
}

func (p CronJob) DoJobs(jobName, jobEvery string, jobHours []string, currentJob jobFunction) {
	for _, jobHour := range jobHours {
		golibs.Log.Printf("| CronJob | Set Job: %s, Every: %s, At: %s\n", jobName, jobEvery, jobHour)
		switch jobEvery {
		case "day":
			gocron.Every(1).Day().At(jobHour).Do(func() { currentJob("CronJob") })
		case "week":
			gocron.Every(1).Week().At(jobHour).Do(func() { currentJob("CronJob") })
		case "month":
			gocron.Every(4).Weeks().At(jobHour).Do(func() { currentJob("CronJob") })
		}
	}
}
