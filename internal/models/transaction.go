package models

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	Id int64 `json:"id,omitempty"`
	DepositId int64 `json:"depositId,omitempty"`
	OwnerUuid uuid.UUID `json:"ownerUuid,omitempty"`
	Amount int64 `json:"amount"`
	Reason string `json:"reason"`
	PartnerUuid uuid.NullUUID `json:"initiatorUserUuid,omitempty"`
	TransactionDate time.Time `json:"transactionDate,omitempty"`
}

type TransactionRepository interface {
	AddTransaction(t *Transaction) (e error)
	//TO DO getTransactions
}

type TransactionUsecase interface {
	CreateTransaction(ownerUuid string, t *Transaction) error
	//TO DO getTransactions
}