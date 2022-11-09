package transactions

import (
	"net/http"

	"github.com/fahmyabdul/golibs"
	"github.com/fahmyabdul/self-growth/data-app/internal/models/transactions"
	"github.com/gin-gonic/gin"
)

// TransactionUpdate godoc
// @Summary      Update Transaction
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        transactionhash   path      string  true  "Transaction Hash"
// @Param 		 request body transactions.Transactions true "Transaction Data"
// @Success      200  {object}  transactions.Transactions
// @Failure		 400  {object}	models.ResponseRestApi
// @Failure		 401  {object}	models.ResponseRestApi
// @Router       /transactions/update/{transactionhash} [patch]
func (p *CtrlTransactions) Update(c *gin.Context) {
	var (
		requestJSON     transactions.Transactions
		transactionHash = c.Param("transactionhash")
	)

	if err := c.ShouldBindJSON(&requestJSON); err != nil {
		errorMsg := err.Error()
		if errorMsg == "EOF" {
			errorMsg = "Request Body must not empty"
		}
		golibs.Log.Printf("| Transactions | Update | Parse Request Body | Error : %s\n", errorMsg)

		c.JSON(http.StatusBadRequest, golibs.ResponseJSON{
			Success: golibs.BoolPointer(false),
			Code:    http.StatusBadRequest,
			Message: errorMsg,
			Data:    nil,
		})

		return
	}

	outputData, httpStatus, err := requestJSON.Update(transactionHash)
	if err != nil {
		golibs.Log.Printf("| Transactions | Update | Request: %v | Error : %s\n", requestJSON, err.Error())

		c.JSON(httpStatus, golibs.ResponseJSON{
			Success: golibs.BoolPointer(false),
			Code:    httpStatus,
			Message: err.Error(),
			Data:    nil,
		})

		return
	}

	c.JSON(http.StatusOK, golibs.ResponseJSON{
		Success: golibs.BoolPointer(true),
		Code:    http.StatusOK,
		Message: "Success",
		Data:    &outputData,
	})
}
