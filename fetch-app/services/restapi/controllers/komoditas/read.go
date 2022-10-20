package komoditas

import (
	"net/http"

	"github.com/fahmyabdul/efishery-task/fetch-app/internal/models"
	"github.com/fahmyabdul/efishery-task/fetch-app/internal/models/komoditas"
	"github.com/fahmyabdul/golibs"
	"github.com/gin-gonic/gin"
)

// KomoditasGetAll godoc
// @Summary      Get All Komoditas Data
// @Tags         komoditas
// @Accept       json
// @Produce      json
// @Success      200  {object}  []komoditas.Komoditas
// @Failure		 400  {object}	models.ResponseRestApi
// @Failure		 401  {object}	models.ResponseRestApi
// @Router       /komoditas/get [get]
// @Security JWT
func (p *CtrlKomoditas) GetAll(c *gin.Context) {
	var usedModels komoditas.Komoditas

	listKomoditas, err := usedModels.GetAll()
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
		Data:   &listKomoditas,
	})
}

// KomoditasGetAggregate godoc
// @Summary      Get The Aggregate Of Komoditas Data
// @Tags         komoditas
// @Accept       json
// @Produce      json
// @Success      200  {object}  []komoditas.KomoditasAggregateContent
// @Failure		 400  {object}	models.ResponseRestApi
// @Failure		 401  {object}	models.ResponseRestApi
// @Router       /komoditas/get/aggregate [get]
// @Security JWT
func (p *CtrlKomoditas) GetAggregate(c *gin.Context) {
	var usedModels komoditas.Komoditas

	listKomoditasAggregate, err := usedModels.GetAggregate()
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
		Data:   &listKomoditasAggregate,
	})
}
