package errors

import "net/http"

type Kind string

type Error struct {
	Kind Kind
	Err error
}

func (e Error) Error() string {
	return string(e.Kind)
}

func E(args ...interface{}) error {
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Kind:
			e.Kind = arg
		case error:
			e.Err = arg
		default:
			panic("unknown behaviour while constructing Error struct")
		}
	}

	return e
}

func GetKind(err error) Kind {
	e, ok := err.(*Error)
	if !ok {
		return UnexpectedErr
	}

	return e.Kind
}

func GetHttpError(err error) int {
	notFoundErrors := []Kind{
		InitiatorDepositNotFoundErr,
		OwnerDepositNotFoundErr,
		RepositoryNoRows,
	}

	badRequestErrors := []Kind {
		FundsValidationErr,
		ReasonValidationErr,
		UuidValidationErr,
		InitiatorFromTargetWithdrawErr,
	}

	paymentRequired := []Kind {
		NotEnoughFundsInitiatorErr,
		NotEnoughFundsOwnerErr,
	}

	internalError := []Kind {
		DepositCreationErr,
		RepositoryTransactionsErr,
		RepositoryDepositsErr,
		RepositoryDownErr,
		RepositoryQueryErr,
		UnexpectedErr,
	}

	kind := GetKind(err)

	if contains(notFoundErrors, kind) {
		return http.StatusNotFound
	}

	if contains(badRequestErrors, kind) {
		return http.StatusBadRequest
	}

	if contains(internalError, kind) {
		return http.StatusInternalServerError
	}

	if contains(paymentRequired, kind) {
		return http.StatusPaymentRequired
	}

	return http.StatusInternalServerError
}

func contains(s []Kind, k Kind) bool {
	isMember := false
	for _, v := range s {
		if k == v {
			isMember = true
		}
	}

	return isMember
}