package main

import (
	"database/sql"

	api "serhatbxld/arf-case/api"
	util "serhatbxld/arf-case/util"

	"github.com/rs/zerolog/log"

	db "serhatbxld/arf-case/db/sqlc"

	_ "github.com/lib/pq"
)

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	dbSource := `postgresql://` + config.DBUser + `:` + config.DBPassword + `@localhost:5432/` + config.DBName + `?sslmode=disable`
	conn, err := sql.Open("postgres", dbSource)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect db")
	}

	store := db.NewStore(conn)

	runGinServer(config, store)
}
