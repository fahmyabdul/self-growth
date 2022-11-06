package middlewares

import (
	"strings"

	"github.com/fahmyabdul/self-growth/fetch-app/services/prometheus"
	"github.com/gin-gonic/gin"
)

func EndpointCounter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if prometheus.EndpointHitsCounter != nil && !strings.Contains(c.Request.URL.Path, "/metrics") {
			prometheus.EndpointHitsCounter.WithLabelValues(c.Request.URL.Path).Inc()
		}
	}
}
