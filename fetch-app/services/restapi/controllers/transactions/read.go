package transactions

import (
	"net/http"

	"github.com/fahmyabdul/golibs"
	"github.com/fahmyabdul/self-growth/fetch-app/internal/models"
	"github.com/fahmyabdul/self-growth/fetch-app/internal/models/transactions"
	"github.com/gin-gonic/gin"
)

// TransactionsGetAll godoc
// @Summary      Get All Transactions Data
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Success      200  {object}  []transactions.Transactions
// @Failure		 400  {object}	models.ResponseRestApi
// @Failure		 401  {object}	models.ResponseRestApi
// @Router       /transactions/get [get]
// @Security JWT
func (p *CtrlTransactions) GetAll(c *gin.Context) {
	var usedModels transactions.Transactions

	listTransactions, err := usedModels.GetAll()
	if err != nil {
		golibs.Log.Printf("| Error : %s\n", err.Error())

		c.JSON(http.StatusBadRequest, models.ResponseRestApi{
			Status: "Error",
			Code:   http.StatusBadRequest,
			Data:   "",
		})

		return
	}

	c.JSON(http.StatusOK, models.ResponseRestApi{
		Status: "Success",
		Code:   http.StatusOK,
		Data:   &listTransactions,
	})
}

// TransactionsGetAggregate godoc
// @Summary      Get The Aggregate Of Transactions Data
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Success      200  {object}  []transactions.TransactionsAggregateContent
// @Failure		 400  {object}	models.ResponseRestApi
// @Failure		 401  {object}	models.ResponseRestApi
// @Router       /transactions/get/aggregate [get]
// @Security JWT
func (p *CtrlTransactions) GetAggregate(c *gin.Context) {
	var usedModels transactions.Transactions

	listTransactionsAggregate, err := usedModels.GetAggregate()
	if err != nil {
		golibs.Log.Printf("| Error : %s\n", err.Error())

		c.JSON(http.StatusBadRequest, models.ResponseRestApi{
			Status: "Error",
			Code:   http.StatusBadRequest,
			Data:   "",
		})

		return
	}

	c.JSON(http.StatusOK, models.ResponseRestApi{
		Status: "Success",
		Code:   http.StatusOK,
		Data:   &listTransactionsAggregate,
	})
}
