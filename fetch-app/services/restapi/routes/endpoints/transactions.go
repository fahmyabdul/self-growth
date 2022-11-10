package endpoints

import (
	"github.com/fahmyabdul/self-growth/fetch-app/services/restapi/controllers/transactions"
	"github.com/fahmyabdul/self-growth/fetch-app/services/restapi/middlewares"
)

func (p *Endpoints) TransactionsEndpoints() {
	transactionsGroup := p.MainGroup.Group("/transactions")
	{
		transactionsGroup.Use(middlewares.JwtAuth(p.BasePath))
		// Get All
		transactionsGroup.GET("/get", (&transactions.CtrlTransactions{}).GetAll)
		// Get All
		transactionsGroup.GET("/get/aggregate", (&transactions.CtrlTransactions{}).GetAggregate)
	}
}
