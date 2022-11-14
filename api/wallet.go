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

// @BasePath /

// ArfCase godoc
// @Summary Create wallet
// @Schemes
// @Description Create wallet for the user
// @Tags wallet
// @Accept json
// @Param Authorization header string true "With the bearer started"
// @Param request body api.createWalletRequest true "Create wallet params"
// @Produce json
// @Success 201 {object} db.Wallet
// @Router /wallets [post]
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

type listWalletRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// @BasePath /

// ArfCase godoc
// @Summary List wallets
// @Schemes
// @Description List wallets of the user
// @Tags wallet
// @Accept json
// @Param Authorization header string true "With the bearer started"
// @Param page_id query integer false "Page Id"
// @Param page_size query integer false "Page Size"
// @Produce json
// @Success 200 {array} arfcasesqlc.Wallet
// @Router /wallets [get]
func (server *Server) listWallets(ctx *gin.Context) {
	var req listWalletRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := server.store.GetUserByUsername(ctx, authPayload.Username)
	arg := db.ListWalletsParams{
		UserID: user.ID,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListWallets(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
