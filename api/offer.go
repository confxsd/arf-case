package api

import (
	"net/http"
	db "serhatbxld/arf-case/db/sqlc"
	"serhatbxld/arf-case/token"

	"github.com/gin-gonic/gin"
)

type offerRequest struct {
	FromCurrency string  `json:"from_currency" binding:"required,min=1"`
	ToCurrency   string  `json:"to_currency" binding:"required,min=1"`
	Rate         float64 `json:"rate" binding:"required,gt=0"`
	Amount       float64 `json:"amount" binding:"required,gt=0"`
}

func (server *Server) createOffer(ctx *gin.Context) {
	var req offerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := server.store.GetUserByUsername(ctx, authPayload.Username)

	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
	}

	arg := db.CreateOfferParams{
		UserID:       user.ID,
		FromCurrency: req.FromCurrency,
		ToCurrency:   req.ToCurrency,
		Rate:         req.Rate,
		Amount:       req.Amount,
	}

	result, err := server.store.CreateOffer(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}
