package api

import (
	"confxsd/arf-case/util"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"

	db "confxsd/arf-case/db/sqlc"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
