package main

import (
	"context"
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

func createSystemUser(ctx context.Context, config util.Config, store db.Store) error {
	arg := db.CreateUserParams{
		Username: config.SystemUsername,
		Password: config.SystemPassword,
	}

	gotUser, err := store.GetUserByUsername(ctx, config.SystemUsername)
	if err != nil {
		user, err := store.CreateUser(ctx, arg)

		if err != nil {
			log.Fatal().Err(err).Msg("cannot create system user")
			return err
		}

		log.Info().Msg("system user created succesfully")

		wallets := []db.CreateWalletParams{
			{
				UserID:   user.ID,
				Balance:  100000,
				Currency: util.USD,
			},
			{
				UserID:   user.ID,
				Balance:  100000,
				Currency: util.EUR,
			},
			{
				UserID:   user.ID,
				Balance:  100000,
				Currency: util.TRY,
			},
		}

		for _, w := range wallets {
			_, err := store.CreateWallet(ctx, w)
			if err != nil {
				log.Fatal().Err(err).Msg("cannot create wallet")
			}
		}

		log.Info().Msg("system user wallets created succesfully")

		return nil
	}

	if gotUser.Username != "" {
		log.Info().Msg("system user already created, skipping initialization")
	}

	return nil
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

	// first time initialization, skips safely if system user available
	createSystemUser(context.Background(), config, store)

	runGinServer(config, store)
}
