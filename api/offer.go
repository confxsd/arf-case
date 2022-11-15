package api

import (
	"context"
	"errors"
	"net/http"
	db "serhatbxld/arf-case/db/sqlc"
	"serhatbxld/arf-case/token"
	"serhatbxld/arf-case/util"
	"time"

	"github.com/gin-gonic/gin"
)

type offerRequest struct {
	FromCurrency string  `json:"from_currency" binding:"required,min=1"`
	ToCurrency   string  `json:"to_currency" binding:"required,min=1"`
	Rate         float64 `json:"rate" binding:"required,gt=0"`
	Amount       float64 `json:"amount" binding:"required,gt=0"`
}

// @BasePath /

// ArfCase godoc
// @Summary Create offer
// @Schemes
// @Description Create offer to convert currencies
// @Tags offer
// @Accept json
// @Param Authorization header string true "With the bearer started"
// @Param request body api.offerRequest true "Create offer params"
// @Produce json
// @Success 201 {object} arfcasesqlc.Offer
// @Router /offers [post]
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
		return
	}

	arg := db.CreateOfferParams{
		UserID:       user.ID,
		FromCurrency: req.FromCurrency,
		ToCurrency:   req.ToCurrency,
		Rate:         req.Rate,
		Amount:       req.Amount,
	}

	offerErr := checkOfferValid(ctx, arg, server.store, server.config)

	if offerErr != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(offerErr))
		return
	}

	result, err := server.store.CreateOffer(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

type approveOfferRequest struct {
	ID int `uri:"id" binding:"required"`
}

func (server *Server) approveOffer(ctx *gin.Context) {
	var request approveOfferRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := server.store.GetUserByUsername(ctx, authPayload.Username)

	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	offer, err := server.store.GetOffer(ctx, int64(request.ID))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if offer.UserID != user.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"msg": "offer not belong to the user"})
		return
	}

	now := time.Now().UTC()

	// offers are valid for 3 minutes
	if now.Sub(offer.CreatedAt) > time.Minute*3 {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "offer timeout"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "voila",
		"id":      request.ID,
	})
}

func determineTransferAmounts(FromCurrency string, ToCurrency string, Amount float64) float64 {
	rate := util.GetRate(FromCurrency, ToCurrency)
	markupRate := util.Markup()

	requiredAmountToTransferFromUser := Amount / (rate * (1 - markupRate))

	return requiredAmountToTransferFromUser
}

/*
e.g

from: USD
to: TRY
amount: 100
rate: 10
markupRate: 0.1
rate * (1 - markupRate) = 9

system should have at least (100) TRY
user should have at least (100/9) = 11.1 USD
(it's supposed to be 10 USD if no markup applied)
*/
func checkOfferValid(ctx context.Context, arg db.CreateOfferParams, store db.Store, config util.Config) error {
	if arg.Rate != util.GetRate(arg.FromCurrency, arg.ToCurrency) {
		return errors.New("Rate is not same with the system")
	}

	if arg.FromCurrency == arg.ToCurrency {
		return errors.New("From & To currencies same, cannot convert.")
	}

	userWallets, err := store.ListWallets(ctx, db.ListWalletsParams{
		UserID: arg.UserID,
		Limit:  5,
		Offset: 0,
	})

	if err != nil {
		return err
	}

	fromCurrencyWallet := db.Wallet{}
	toCurrencyWallet := db.Wallet{}

	for _, w := range userWallets {
		if w.Currency == arg.FromCurrency {
			fromCurrencyWallet = w
		}
		if w.Currency == arg.ToCurrency {
			toCurrencyWallet = w
		}
	}

	if fromCurrencyWallet.Currency == "" || toCurrencyWallet.Currency == "" {
		return errors.New("User has not required wallets to convert")
	}

	requiredAmountToTransferFromUser := determineTransferAmounts(arg.FromCurrency, arg.ToCurrency, arg.Rate)

	if float64(fromCurrencyWallet.Balance) < requiredAmountToTransferFromUser {
		return errors.New("User has not enough balance to convert")
	}

	return nil
}
