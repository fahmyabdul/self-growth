package transactions

import (
	"io/ioutil"
	"net/http"

	"github.com/fahmyabdul/golibs"
	"github.com/fahmyabdul/self-growth/data-app/internal/models"
	"github.com/fahmyabdul/self-growth/data-app/internal/models/transactions"
	"github.com/gin-gonic/gin"
)

// TransactionsGetAll godoc
// @Summary      Get All Transaction Data
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Success      200  {object}  []transactions.Transactions
// @Failure		 400  {object}	models.ResponseRestApi
// @Failure		 401  {object}	models.ResponseRestApi
// @Router       /transactions/get [get]
func (p *CtrlTransactions) GetAll(c *gin.Context) {
	var usedModels transactions.Transactions

	listTransactions, _, httpStatus, err := usedModels.GetAll()
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
		Code:   httpStatus,
		Data:   &listTransactions,
	})
}

// TransactionsGetByUsername godoc
// @Summary      Get Transaction Data By Username
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        username   path      string  true  "Username"
// @Success      200  {object}  []transactions.Transactions
// @Failure		 400  {object}	models.ResponseRestApi
// @Failure		 401  {object}	models.ResponseRestApi
// @Router       /transactions/get/username/{username} [get]
func (p *CtrlTransactions) GetByUsername(c *gin.Context) {
	var (
		usedModels transactions.Transactions
		username   = c.Param("username")
	)

	outputData, httpStatus, err := usedModels.GetByUsername(username)
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

// TransactionsGetByFilter godoc
// @Summary      Get Transaction Data By Custom Filter
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param 		 request body map[string]interface{} true "Filter Data"
// @Success      200  {object}  []transactions.Transactions
// @Failure		 400  {object}	models.ResponseRestApi
// @Failure		 401  {object}	models.ResponseRestApi
// @Router       /transactions/get/filter [post]
func (p *CtrlTransactions) GetByFilter(c *gin.Context) {
	var usedModels transactions.Transactions

	requestJSON, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		golibs.Log.Printf("| Error : %s\n", err.Error())

		c.JSON(http.StatusBadRequest, models.ResponseRestApi{
			Status: "Error",
			Code:   http.StatusBadRequest,
			Data:   "",
		})

		return
	}

	outputData, httpStatus, err := usedModels.GetByFilter(requestJSON)
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
