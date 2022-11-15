package api

import (
	"context"
	"errors"
	"net/http"
	db "serhatbxld/arf-case/db/sqlc"
	"serhatbxld/arf-case/token"
	"serhatbxld/arf-case/util"

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

	ctx.JSON(http.StatusOK, result)
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

	rate := util.GetRate(arg.FromCurrency, arg.ToCurrency)
	markupRate := util.Markup()

	requiredAmountToTransferFromUser := arg.Amount / (rate * (1 - markupRate))

	if float64(fromCurrencyWallet.Balance) < requiredAmountToTransferFromUser {
		return errors.New("User has not enough balance to convert")
	}

	systemUser, err := store.GetUserByUsername(ctx, config.SystemUsername)
	if err != nil {
		return err
	}

	systemWallets, err := store.ListWallets(ctx, db.ListWalletsParams{
		UserID: systemUser.ID,
		Limit:  5,
		Offset: 0,
	})

	systemToCurrencyWallet := db.Wallet{}
	for _, w := range systemWallets {
		if w.Currency == arg.ToCurrency {
			systemToCurrencyWallet = w
		}
	}

	if float64(systemToCurrencyWallet.Balance) < arg.Amount {
		return errors.New("System has not enough balance to convert")
	}

	return nil
}
