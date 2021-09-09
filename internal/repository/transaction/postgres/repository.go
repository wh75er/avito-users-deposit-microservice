package postgres

import (
	"bank-microservice/internal/errors"
	"bank-microservice/internal/models"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TransactionRepository struct {
	Db *sqlx.DB
	Logger *logrus.Logger
}

func New(db *sqlx.DB, logger *logrus.Logger) models.TransactionRepository {
	return &TransactionRepository{ db, logger }
}

func (r *TransactionRepository) AddTransaction(t *models.Transaction) (e error) {
	_, e = r.Db.Exec("INSERT INTO Transactions (DepositId, OwnerUuid, Amount, Reason, PartnerUuid, TransactionDate) " +
		"VALUES($1, $2, $3, $4, $5, $6)", t.DepositId, t.OwnerUuid, t.Amount, t.Reason, t.PartnerUuid, t.TransactionDate)
	if e == sql.ErrConnDone {
		r.Logger.Error("Database error: ", e.Error())
		e = errors.E(errors.RepositoryDownErr, e)
	} else if e != nil {
		r.Logger.Error("Database error: ", e.Error())
		e = errors.E(errors.RepositoryQueryErr, e)
	}

	return
}