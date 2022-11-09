package endpoints

import (
	"github.com/fahmyabdul/self-growth/data-app/services/restapi/controllers/transactions"
)

func (p *Endpoints) KomoditasEndpoints() {
	transactionsGroup := p.MainGroup.Group("/transactions")
	{
		// transactionsGroup.Use(middlewares.JwtAuth(p.BasePath))
		// Get All
		transactionsGroup.GET("/get", (&transactions.CtrlTransactions{}).GetAll)
		// Get By Username
		transactionsGroup.GET("/get/username/:username", (&transactions.CtrlTransactions{}).GetByUsername)
		// Get By Filter
		transactionsGroup.POST("/get/filter", (&transactions.CtrlTransactions{}).GetByFilter)
		// Create
		transactionsGroup.POST("/create", (&transactions.CtrlTransactions{}).Create)
		// Delete
		transactionsGroup.DELETE("/delete/:transactionhash", (&transactions.CtrlTransactions{}).Delete)
		// Update
		transactionsGroup.PATCH("/update/:transactionhash", (&transactions.CtrlTransactions{}).Update)
	}
}
