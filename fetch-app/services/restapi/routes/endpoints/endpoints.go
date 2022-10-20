package endpoints

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/fahmyabdul/efishery-task/fetch-app/services/restapi/controllers"
)

type Endpoints struct {
	Router    *gin.Engine
	BasePath  string
	MainGroup *gin.RouterGroup
}

func New(routes *gin.Engine, basePath string) (*Endpoints, error) {
	if routes == nil {
		return nil, fmt.Errorf("routes must not null")
	}
	return &Endpoints{Router: routes, BasePath: basePath}, nil
}

// LoadEndpoints : Function for all possible route
func (p *Endpoints) LoadEndpoints() {
	// Main Router Group
	p.MainGroup = p.Router.Group("/").Group(p.BasePath)
	{
		p.MainGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		p.MainGroup.GET("/heartbeat", (&controllers.Heartbeat{}).GetVersion)
		p.MainGroup.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}

	// Set Komoditas Endpoints
	p.KomoditasEndpoints()
}
