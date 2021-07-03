package komoditas

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/risyard/efishery-task/fetch-app/logic/komoditas"
	"github.com/risyard/efishery-task/fetch-app/model"
)

type IKomoditasHandler interface {
	GetListKomoditas(ctx *gin.Context)
}

type KomoditasHandler struct {
	KomLogic komoditas.IKomoditasLogic
}

func NewKomoditasHandler() IKomoditasHandler {
	return &KomoditasHandler{
		KomLogic: komoditas.NewKomoditasLogic(),
	}
}

func (h *KomoditasHandler) GetListKomoditas(ctx *gin.Context) {
	listKomoditas, err := h.KomLogic.GetListKomoditas()
	if err != nil {
		ctx.JSON(500, model.BadResponse{
			Status:  500,
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Status: 200,
		Data:   listKomoditas,
	})
}
