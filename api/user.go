package api

import (
	"net/http"
	db "serhatbxld/arf-case/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type userResponse struct {
	Username string `json:"username"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username: user.Username,
	}
}

// @BasePath /

// ArfCase godoc
// @Summary Create user
// @Schemes
// @Description Create user with username & password
// @Tags user
// @Accept json
// @Param request body api.createUserRequest true "Create user params"
// @Produce json
// @Success 201 {object} api.userResponse
// @Router /users [post]
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Password: req.Password,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}
