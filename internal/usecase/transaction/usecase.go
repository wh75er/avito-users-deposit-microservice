package transaction

import (
	"bank-microservice/internal/errors"
	"bank-microservice/internal/models"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Usecase struct {
	DepositRepository models.DepositRepository
	TransactionRepository models.TransactionRepository
	Logger *logrus.Logger
}

func New(DepositRep models.DepositRepository, TransactionRep models.TransactionRepository, logger *logrus.Logger) models.TransactionUsecase {
	return &Usecase{DepositRep, TransactionRep, logger}
}

func (u *Usecase) CreateTransaction(ownerUuid string, t *models.Transaction) error {
	var e error

	if t.Amount == 0 {
		e = errors.E(errors.FundsValidationErr)
		u.Logger.Error("Usecase error: ", e.Error())
		return e
	}

	if t.Reason == "" || len(t.Reason) > 250 {
		e = errors.E(errors.ReasonValidationErr)
		u.Logger.Error("Usecase error: ", e.Error())
		return e
	}

	validOwnerUuid, e := uuid.Parse(ownerUuid)
	if e != nil {
		e = errors.E(errors.UuidValidationErr, e)
		u.Logger.Error("Usecase error: ", e.Error())
		return e
	}

	if t.PartnerUuid.Valid {
		e = u.makeDuoTransaction(validOwnerUuid, t.PartnerUuid.UUID, t)
		if e != nil {
			u.Logger.Error("Usecase error: ", e.Error())
			return e
		}
	} else {
		e = u.makeSoloTransaction(validOwnerUuid, t)
		if e != nil {
			u.Logger.Error("Usecase error: ", e.Error())
			return e
		}
	}

	return nil
}
