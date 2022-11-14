package api

import (
	"net/http"
	db "serhatbxld/arf-case/db/sqlc"
	"serhatbxld/arf-case/token"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createWalletRequest struct {
	Currency string `json:"currency" binding:"required,currency"`
}

func (server *Server) createWallet(ctx *gin.Context) {
	var req createWalletRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := server.store.GetUserByUsername(ctx, authPayload.Username)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
	}

	arg := db.CreateWalletParams{
		UserID:   user.ID,
		Currency: req.Currency,
		Balance:  0,
	}

	Wallet, err := server.store.CreateWallet(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, Wallet)
}
