package errors

const (
	DepositCreationErr = Kind("Failed to create deposit")
	FundsValidationErr = Kind("Amount of money must not be 0")
	RepositoryTransactionsErr = Kind("Something wrong with transactions repository")
	RepositoryDepositsErr = Kind("Something wrong with deposit repository")
	RepositoryDownErr = Kind("Repository connection problem")
	RepositoryQueryErr = Kind("Failed to perform query")
	RepositoryNoRows = Kind("No rows were found")
	ReasonValidationErr = Kind("Reason length must contain data and be less than 250 symbols")
	InitiatorDepositNotFoundErr = Kind("Initiator of user to user money transfer doesn't have deposit account")
	InitiatorFromTargetWithdrawErr = Kind("Initiator of user to user money transfer cannot request target's money withdraw")
	NotEnoughFundsInitiatorErr = Kind("Initiator of user to user money transfer doesn't have enough funds")
	NotEnoughFundsOwnerErr = Kind("Deposit owner doesn't have enough funds to perform this operation")
	OwnerDepositNotFoundErr = Kind("User doesn't have a deposit to perform this operation")
	UuidValidationErr = Kind("Provided uuid is not valid")
	UnexpectedErr = Kind("Unexpected error occurred")
)
