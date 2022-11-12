package main

import (
	"database/sql"

	api "serhatbxld/arf-case/api"
	util "serhatbxld/arf-case/util"

	"github.com/rs/zerolog/log"

	db "serhatbxld/arf-case/db/sqlc"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"

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

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	dbSource := `postgresql://` + config.DBUser + `:` + config.DBPassword + `@postgres:` + config.DBPort + `/` + config.DBName + `?sslmode=disable`
	conn, err := sql.Open("postgres", dbSource)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect db")
	}

	runDBMigration(config.MigrationURL, dbSource)

	store := db.NewStore(conn)

	runGinServer(config, store)
}
