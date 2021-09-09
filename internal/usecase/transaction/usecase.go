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
	if t.Amount == 0 {
		return errors.E(errors.FundsValidationErr)
	}

	if t.Reason == "" || len(t.Reason) > 250 {
		return errors.E(errors.ReasonValidationErr)
	}

	validOwnerUuid, e := uuid.Parse(ownerUuid)
	if e != nil {
		return errors.E(errors.UuidValidationErr, e)
	}

	if t.PartnerUuid.Valid {
		e = u.makeDuoTransaction(validOwnerUuid, t.PartnerUuid.UUID, t)
		if e != nil {
			return e
		}
	} else {
		e = u.makeSoloTransaction(validOwnerUuid, t)
		if e != nil {
			return e
		}
	}

	return nil
}
