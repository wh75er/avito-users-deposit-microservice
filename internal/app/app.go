package app

import (
	transactionDelivery "bank-microservice/internal/delivery/transaction/http"
	depositRepPostgres "bank-microservice/internal/repository/deposit/postgres"
	transactionRepPostgres "bank-microservice/internal/repository/transaction/postgres"
	transactionUsecase "bank-microservice/internal/usecase/transaction"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type App struct {
	db *sqlx.DB
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

	depositRepository := depositRepPostgres.New(a.db, a.logger)
	transactionRepository := transactionRepPostgres.New(a.db, a.logger)

	transactionUcase := transactionUsecase.New(depositRepository, transactionRepository, a.logger)

	transactionDelivery.NewTransactionHandler(transactionUcase, a.server, a.logger)

	a.server.Use(middleware.Logger())

	if err := a.server.Start(":" + strconv.Itoa(a.conf.Server.Port)); err == http.ErrServerClosed {
		a.logger.Fatal(err)
	}
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

	a.db = db
}
