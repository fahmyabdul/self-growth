package prometheus

import (
	"github.com/fahmyabdul/efishery-task/fetch-app/app"
	"github.com/fahmyabdul/golibs"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

type Prometheus struct{}

var (
	// AppStatusGauge :
	AppStatusGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace:   "test",
			Subsystem:   "app",
			Name:        "status",
			Help:        "TestService - App Status",
			ConstLabels: prometheus.Labels(map[string]string{"version": app.CurrentVersion}),
		},
	)

	// KafkaStatusGauge :
	KafkaStatusGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "test",
			Subsystem: "kafka",
			Name:      "status",
			Help:      "TestService - Kafka Status",
		},
	)

	// EndpointHitsCounter :
	EndpointHitsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "test",
			Subsystem: "endpoint",
			Name:      "hits",
			Help:      "The total number of endpoints hitted",
		},
		[]string{
			"url",
		},
	)
)

// Start : Starting Prometheus
func (p Prometheus) Start() error {
	golibs.Log.Println("| Prometheus | Initialize")

	err := prometheus.Register(collectors.NewBuildInfoCollector())
	if err != nil {
		return err
	}

	err = prometheus.Register(AppStatusGauge)
	if err != nil {
		return err
	}

	err = prometheus.Register(KafkaStatusGauge)
	if err != nil {
		return err
	}

	err = prometheus.Register(EndpointHitsCounter)
	if err != nil {
		return err
	}

	golibs.Log.Println("| Prometheus | Initialize Done")

	return nil
}
