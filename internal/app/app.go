package app

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type App struct {
	Db *sqlx.DB
	conf *config
	configName string
	server *echo.Echo
	logger *logrus.Logger
}

func New() *App {
	return &App {
		nil,
		newConfig(),
		"",
		echo.New(),
		logrus.New(),
	}
}


func (a *App) Run(configFilename string) {
	a.configName = configFilename
	a.setupApp()
	a.setupStorage()

	// TO DO init all delivery, usecases, repositories
	// TO DO start server
}

func (a *App) setupStorage() {
	a.establishDbConnection(a.conf.Storage.MaxPoolConn, a.conf.Storage.driver, a.conf.Storage.url)
	a.logger.Info("Successfully established connection with database")
}

func (a *App) setupApp() {
	// Check if there is a configuration file
	if a.configName != "" {
		if err := a.conf.loadFromToml(a.configName); err != nil {
			a.logger.Fatal("Failed to decode configuration file: ", err.Error())
		}
		a.logger.Infof("Loaded configuration file: %v. Current configuration: %v", a.configName, a.conf)
	} else {
		a.logger.Warnf("Configuration file is not specified, using default configuration: %v", a.conf)
	}
}

func (a *App) establishDbConnection(maxPoolConn int, driverName, dbUrl string) {
	db, err := sqlx.Open(driverName, dbUrl)
	if err != nil {
		a.logger.Fatal("Failed to establish connection with db: ", err.Error())
	}

	err = db.Ping()
	if err != nil {
		a.logger.Fatal("Failed to ping db ", err.Error())
	}

	db.SetMaxOpenConns(maxPoolConn)

	a.Db = db
}
