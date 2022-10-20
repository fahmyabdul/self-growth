package endpoints

import (
	"github.com/fahmyabdul/efishery-task/fetch-app/services/restapi/controllers/komoditas"
	"github.com/fahmyabdul/efishery-task/fetch-app/services/restapi/middlewares"
)

func (p *Endpoints) KomoditasEndpoints() {
	komoditasGroup := p.MainGroup.Group("/komoditas")
	{
		komoditasGroup.Use(middlewares.JwtAuth(p.BasePath))
		// Get All
		komoditasGroup.GET("/get", (&komoditas.CtrlKomoditas{}).GetAll)
		// Get All
		komoditasGroup.GET("/get/aggregate", (&komoditas.CtrlKomoditas{}).GetAggregate)
	}
}
