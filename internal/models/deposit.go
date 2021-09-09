package models

import (
	"github.com/google/uuid"
	"time"
)

type Deposit struct {
	Id int64 `json:",omitempty"`
	UserUuid uuid.UUID
	Deposit int64
	CreationDate time.Time
}

type DepositRepository interface {
	GetDepositByOwner(ownerUuid uuid.UUID) (d Deposit, e error)
	AddNewDepositForOwner(d *Deposit) (id int64, e error)
	UpdateDepositByOwner(d *Deposit) (e error)
}

type DepositUsecase interface {
	GetDepositByOwner(ownerUuid string) (Deposit, error)
}