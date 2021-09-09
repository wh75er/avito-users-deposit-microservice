package deposit

import (
	"bank-microservice/internal/errors"
	"bank-microservice/internal/models"
	ucaseHelpers "bank-microservice/internal/usecase"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Usecase struct {
	DepositRepository models.DepositRepository
	Logger *logrus.Logger
}

func New(DepositRep models.DepositRepository, logger *logrus.Logger) models.DepositUsecase {
	return &Usecase{DepositRep, logger}
}

func (u *Usecase) GetDepositByOwner(ownerUuid string) (d models.Deposit, e error) {
	validOwnerUuid, e := uuid.Parse(ownerUuid)
	if e != nil {
		e = errors.E(errors.UuidValidationErr, e)
		u.Logger.Error("Usecase error: ", e.Error())
		return
	}

	d, depositExists, e := ucaseHelpers.GetUsersDepositByUuid(u.DepositRepository, validOwnerUuid)
	if e != nil {
		u.Logger.Error("Usecase error: ", e.Error())
		return
	}

	if !depositExists {
		e = errors.E(errors.OwnerDepositNotFoundErr)
		u.Logger.Error("Usecase error: ", e.Error())
		return
	}

	return
}