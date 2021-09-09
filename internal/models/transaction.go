package models

import (
	"github.com/google/uuid"
)

type Transaction struct {
	Id int64 `json:",omitempty"`
	DepositId int64 `json:",omitempty"`
	OwnerUuid uuid.UUID
	Amount int64
	Reason string
	PartnerUuid uuid.NullUUID `json:"initiatorUserUuid,omitempty"`
	TransactionDate int64
}

type TransactionRepository interface {
	AddTransaction(t *Transaction) (e error)
	//TO DO getTransactions
}

type TransactionUsecase interface {
	CreateTransaction(ownerUuid string, t *Transaction) error
	//TO DO getTransactions
}