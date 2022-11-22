package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	db "confxsd/arf-case/db/sqlc"
	docs "confxsd/arf-case/docs"
	"confxsd/arf-case/token"
	util "confxsd/arf-case/util"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

// @BasePath /

// ArfCase godoc
// @Summary index sample
// @Schemes
// @Description do test
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} voila
// @Router / [get]
func getting(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "voila",
	})
}

func protected(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "protected voila",
	})
}

func (server *Server) setupRouter() {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/"

	// public routes
	r.GET("/", getting)
	r.POST("/users", server.createUser)
	r.POST("/auth", server.auth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// protected routes
	authRoutes := r.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/me", protected) // to test :D

	authRoutes.POST("/wallets", server.createWallet)
	authRoutes.GET("/wallets", server.listWallets)

	authRoutes.POST("/offers", server.createOffer)
	authRoutes.POST("/offers/:id/approve", server.approveOffer)

	server.router = r
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
