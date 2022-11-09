package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	docs "serhatbxld/arf-case/docs"
	util "serhatbxld/arf-case/util"

	"github.com/rs/zerolog/log"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

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

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", getting)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	fmt.Print(config)

	r := setupRouter()
	docs.SwaggerInfo.BasePath = "/"

	r.Run() // listen and serve on 8080
}
