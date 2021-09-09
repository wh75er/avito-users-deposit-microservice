package transaction

import (
	"bank-microservice/internal/errors"
	"bank-microservice/internal/models"
	ucaseHelpers "bank-microservice/internal/usecase"
	"github.com/google/uuid"
	"time"
)

func (u *Usecase) makeSoloTransaction(ownerUuid uuid.UUID, t *models.Transaction) error {
	// Check if owner's deposit exists
	d, depositExists, e := ucaseHelpers.GetUsersDepositByUuid(u.DepositRepository, ownerUuid)
	if e != nil {
		return e
	}

	// If not exists and money withdrawing operation performed
	// exit with error
	if !depositExists && t.Amount < 0 {
		e = errors.E(errors.OwnerDepositNotFoundErr)
		return e
	}

	// if deposit not exists, but there is money raising operation
	// init and create deposit for the user with 0 balance value
	if !depositExists && t.Amount > 0 {
		d, e = u.initDepositForUser(ownerUuid)
		if e != nil {
			return e
		}
	}

	// Change deposit balance
	d.Deposit, e = u.changeDepositBalance(d.Deposit, t.Amount)
	if e != nil {
		return e
	}

	// Init transaction fields
	t.OwnerUuid = ownerUuid
	t.DepositId = d.Id
	t.TransactionDate = time.Now().UTC()

	// Make transaction
	e = u.TransactionRepository.AddTransaction(t)
	if e != nil {
		e = errors.E(errors.RepositoryTransactionsErr, e)
		return e
	}

	// Update deposit balance
	e = u.DepositRepository.UpdateDepositByOwner(&d)
	if e != nil {
		e = errors.E(errors.RepositoryDepositsErr)
		return e
	}

	return nil
}

func (u *Usecase) makeDuoTransaction(targetUuid uuid.UUID, initiatorUuid uuid.UUID, t *models.Transaction) error {
	var e error

	// if operation is money withdraw - exit
	if t.Amount < 0 {
		e = errors.E(errors.InitiatorFromTargetWithdrawErr)
		return e
	}

	// Find initiator's deposit
	initiatorDeposit, initiatorDepositExists, e := ucaseHelpers.GetUsersDepositByUuid(u.DepositRepository, initiatorUuid)
	if e != nil {
		return e
	}

	// If initiator doesn't have a deposit - exit
	if !initiatorDepositExists {
		e = errors.E(errors.InitiatorDepositNotFoundErr)
		return e
	}

	// Find target's deposit
	targetDeposit, targetDepositExists, e := ucaseHelpers.GetUsersDepositByUuid(u.DepositRepository, targetUuid)
	if e != nil {
		return e
	}

	// if Target's deposit doesn't exist and operation is funds raising
	// init deposit with 0 balance
	if !targetDepositExists && t.Amount > 0 {
		if !targetDepositExists && t.Amount > 0 {
			targetDeposit, e = u.initDepositForUser(targetUuid)
			if e != nil {
				return e
			}
		}
	}

	// Change initiator balance, if not enough funds - exit
	initiatorDeposit.Deposit, e = u.changeDepositBalance(initiatorDeposit.Deposit, -t.Amount)
	if e != nil {
		e = errors.E(errors.NotEnoughFundsInitiatorErr, e)
		return e
	}

	// Change target balance, if error - unexpected exit(t.Amount should be >= 0)
	targetDeposit.Deposit, e = u.changeDepositBalance(targetDeposit.Deposit, t.Amount)
	if e != nil {
		e = errors.E(errors.UnexpectedErr, e)
		return e
	}

	// Change initiator deposit balance in repository
	e = u.DepositRepository.UpdateDepositByOwner(&initiatorDeposit)
	if e != nil {
		e = errors.E(errors.RepositoryDepositsErr)
		return e
	}

	// Change target deposit balance in repository
	e = u.DepositRepository.UpdateDepositByOwner(&targetDeposit)
	if e != nil {
		e = errors.E(errors.RepositoryDepositsErr)
		return e
	}

	currentTime := time.Now().UTC()

	// Init transaction for target
	targetTransaction := models.Transaction {
		DepositId: targetDeposit.Id,
		OwnerUuid: targetUuid,
		Amount: t.Amount,
		Reason: t.Reason,
		PartnerUuid: uuid.NullUUID { UUID: initiatorUuid, Valid: true },
		TransactionDate: currentTime,
	}

	// Make transaction for target
	e = u.TransactionRepository.AddTransaction(&targetTransaction)
	if e != nil {
		e = errors.E(errors.RepositoryTransactionsErr, e)
		return e
	}

	// Init transaction for initiator
	initiatorTransaction := models.Transaction {
		DepositId: initiatorDeposit.Id,
		OwnerUuid: initiatorUuid,
		Amount: -t.Amount,
		Reason: t.Reason,
		PartnerUuid: uuid.NullUUID { UUID: targetUuid, Valid: true},
		TransactionDate: currentTime,
	}

	// Make transaction for initiator
	e = u.TransactionRepository.AddTransaction(&initiatorTransaction)
	if e != nil {
		e = errors.E(errors.RepositoryTransactionsErr, e)
		return e
	}

	return nil
}

func (u *Usecase) initDepositForUser(targetUuid uuid.UUID) (d models.Deposit, e error) {
	d = models.Deposit{UserUuid: targetUuid, Deposit: 0, CreationDate: time.Now().UTC()}

	id, e := u.DepositRepository.AddNewDepositForOwner(&d)
	if e != nil {
		e = errors.E(errors.DepositCreationErr, e)
		return
	}
	d.Id = id

	return
}

func (u *Usecase) changeDepositBalance(balance int64, delta int64) (newBalance int64, e error) {
	if newBalance = balance + delta; newBalance < 0 {
		e = errors.E(errors.NotEnoughFundsOwnerErr)
	}

	return
}