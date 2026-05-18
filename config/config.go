package config

import (
	"crypto/tls"
	"database/sql"
	"log/slog"
	logger "serv-test/log"

	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
)

type Application struct {
	Logger         *slog.Logger
	DB             *sql.DB
	SessionManager *scs.SessionManager
	FormDecoder    *form.Decoder
}

var App = &Application{
	Logger: logger.Logger,
	// DB is initialized in database.go DatabaseConnect function
	SessionManager: scs.New(),
	FormDecoder:    form.NewDecoder(),
}

var TlsConfig = &tls.Config{
	CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
}
