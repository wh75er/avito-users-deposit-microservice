CREATE TABLE IF NOT EXISTS Deposits(
    Id serial PRIMARY KEY,
    UserUuid UUID NOT NULL UNIQUE,
    Deposit INT NOT NULL,
    CreationDate TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS Transactions(
    Id serial PRIMARY KEY,
    DepositId INT REFERENCES Deposits(Id),
    OwnerUuid UUID NOT NULL,
    Amount INT NOT NULL,
    Reason VARCHAR(250) NOT NULL,
    PartnerUuid UUID,
    TransactionDate TIMESTAMP NOT NULL
);
