package models

import (
	"github.com/google/uuid"
	"time"
)

type Deposit struct {
	Id int64 `json:"omitempty"`
	UserUuid uuid.UUID
	Deposit int64
	CreationDate time.Time
}

type DepositRepository interface {
	GetDepositIdByOwner(ownerUuid uuid.UUID) (id int64, e error)
	GetDepositByOwner(ownerUuid uuid.UUID) (d Deposit, e error)
	GetDepositById(id int64) (d *Deposit, e error)
	AddNewDepositForOwner(d *Deposit) (id int64, e error)
	UpdateDepositByOwner(d *Deposit) (e error)
}

type DepositUsecase interface {
	getDepositByOwner(ownerUuid uuid.UUID) (Deposit, error)
}