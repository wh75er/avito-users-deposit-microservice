package models

import (
	"github.com/google/uuid"
	"time"
)

type Deposit struct {
	Id int64 `json:"id,omitempty"`
	UserUuid uuid.UUID `json:"userUuid"`
	Deposit int64 `json:"deposit"`
	CreationDate time.Time `json:"creationDate"`
}

type DepositRepository interface {
	GetDepositByOwner(ownerUuid uuid.UUID) (d Deposit, e error)
	AddNewDepositForOwner(d *Deposit) (id int64, e error)
	UpdateDepositByOwner(d *Deposit) (e error)
}

type DepositUsecase interface {
	GetDepositByOwner(ownerUuid string) (Deposit, error)
}