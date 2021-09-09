package models

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	Id int64 `json:",omitempty"`
	DepositId int64 `json:",omitempty"`
	OwnerUuid uuid.UUID `json:",omitempty"`
	Amount int64
	Reason string
	PartnerUuid uuid.NullUUID `json:"initiatorUserUuid,omitempty"`
	TransactionDate time.Time `json:",omitempty"`
}

type TransactionRepository interface {
	AddTransaction(t *Transaction) (e error)
	//TO DO getTransactions
}

type TransactionUsecase interface {
	CreateTransaction(ownerUuid string, t *Transaction) error
	//TO DO getTransactions
}