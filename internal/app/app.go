package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"log"
)

type App struct {
	Db *sqlx.DB
	conf *config
	server *echo.Echo
}

func (a *App) ConfigureApp() {
	a.conf = &config{}
	a.conf.loadFromToml("release.toml")
}

func (a *App) establishDbConnection(maxPoolConn int, dbUrl string) {
	db, err := sqlx.Open("pgx", dbUrl)
	if err != nil {
		log.Fatal("failed to establish connection with db: " + err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to ping db " + err.Error())
	}

	db.SetMaxOpenConns(10)

	a.Db = db
}
