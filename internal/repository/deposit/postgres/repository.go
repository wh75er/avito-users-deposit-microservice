package postgres

import (
	"bank-microservice/internal/errors"
	"bank-microservice/internal/models"
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type DepositRepository struct {
	Db *sqlx.DB
	Logger *logrus.Logger
}

func New(db *sqlx.DB, logger *logrus.Logger) models.DepositRepository {
	return &DepositRepository{ db, logger }
}

func (r *DepositRepository) GetDepositByOwner(ownerUuid uuid.UUID) (d models.Deposit, e error) {
	e = r.Db.Get(&d, "SELECT id, userUuid, deposit, creationDate FROM deposits WHERE userUuid = $1", ownerUuid)
	if e == sql.ErrConnDone {
		r.Logger.Error("Database error: ", e.Error())
		e = errors.E(errors.RepositoryDownErr, e)
	} else if e == sql.ErrNoRows {
		r.Logger.Error("Database error: ", e.Error())
		e = errors.E(errors.RepositoryNoRows, e)
	} else if e != nil {
		r.Logger.Error("Database error: ", e.Error())
		e = errors.E(errors.RepositoryQueryErr)
	}

	return
}

func (r *DepositRepository) AddNewDepositForOwner(d *models.Deposit) (id int64, e error) {
	e = r.Db.QueryRowx("INSERT INTO deposits(userUuid, deposit, creationDate) VALUES ($1, $2, $3) RETURNING id",
		d.UserUuid, d.Deposit, d.CreationDate).Scan(&id)
	if e == sql.ErrConnDone {
		r.Logger.Error("Database error: ", e.Error())
		e = errors.E(errors.RepositoryDownErr, e)
	} else if e != nil {
		r.Logger.Error("Database error: ", e.Error())
		e = errors.E(errors.RepositoryQueryErr, e)
	}

	return
}

func (r *DepositRepository) UpdateDepositByOwner(d *models.Deposit) (e error) {
	_, e = r.Db.Exec("UPDATE deposits SET deposit = $1 WHERE userUuid = $2",
		d.Deposit, d.UserUuid)
	if e == sql.ErrConnDone {
		r.Logger.Error("Database error: ", e.Error())
		e = errors.E(errors.RepositoryDownErr, e)
	} else if e != nil {
		r.Logger.Error("Database error: ", e.Error())
		e = errors.E(errors.RepositoryQueryErr, e)
	}

	return
}
