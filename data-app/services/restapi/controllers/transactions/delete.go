package transactions

import (
	"net/http"

	"github.com/fahmyabdul/golibs"
	"github.com/fahmyabdul/self-growth/data-app/internal/models"
	"github.com/fahmyabdul/self-growth/data-app/internal/models/transactions"
	"github.com/gin-gonic/gin"
)

// TransactionsDelete godoc
// @Summary      Delete Transaction Data By Transaction Hash
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        transactionhash   path      string  true  "Transaction Hash"
// @Success      200  {object}  models.ResponseRestApi
// @Failure		 400  {object}	models.ResponseRestApi
// @Failure		 401  {object}	models.ResponseRestApi
// @Router       /transactions/delete/{transactionhash} [delete]
func (p *CtrlTransactions) Delete(c *gin.Context) {
	var (
		usedModels      transactions.Transactions
		transactionHash = c.Param("transactionhash")
	)

	outputData, httpStatus, err := usedModels.Delete(transactionHash)
	if err != nil {
		golibs.Log.Printf("| Error : %s\n", err.Error())

		c.JSON(httpStatus, models.ResponseRestApi{
			Status: "Error",
			Code:   httpStatus,
			Data:   "",
		})

		return
	}

	c.JSON(http.StatusOK, models.ResponseRestApi{
		Status: "Success",
		Code:   http.StatusOK,
		Data:   &outputData,
	})
}
