package main

import (
	"database/sql"
	"gfg/pkg/api"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	db, err := sql.Open("mysql", "user:password@tcp(db:3306)/product")

	if err != nil {
		log.Error().Err(err).Msg("Fail to create server")
		return
	}

	defer db.Close()

	activateSMSProvider := os.Getenv("SMS_PROVIDER") == "true"
	activateEmailProvider := os.Getenv("EMAIL_PROVIDER") == "true"

	apiHost := os.Getenv("API_HOST")
	if apiHost == "" {
		panic("API_HOST env var is missing")
	}

	apiScheme := os.Getenv("API_SCHEME")
	if apiHost == "" {
		panic("API_SCHEME env var is missing")
	}

	engine, err := api.CreateAPIEngine(db, activateSMSProvider, activateEmailProvider, apiHost, apiScheme)

	if err != nil {
		log.Error().Err(err).Msg("Fail to create server")
		return
	}

	log.Info().Msg("Start server")
	log.Fatal().Err(engine.Run(os.Getenv("LISTEN"))).Msg("Fail to listen and serve")
}
