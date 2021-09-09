package usecase

import (
	"bank-microservice/internal/errors"
	"bank-microservice/internal/models"
	"github.com/google/uuid"
)


func GetUsersDepositByUuid(depositRepository models.DepositRepository, targetUuid uuid.UUID) (d models.Deposit, exists bool, e error) {
	exists = true

	// Check if owner deposit exists
	d, e = depositRepository.GetDepositByOwner(targetUuid)
	if e != nil {
		if errors.GetKind(e) == errors.RepositoryNoRows {
			e = nil
			exists = false
		} else {
			e = errors.E(errors.RepositoryDepositsErr, e)
			return
		}
	}

	return
}
