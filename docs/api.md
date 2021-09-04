# Markdown API specification

Simple and small markdown api documentation

## Contents:

List of contents:

- [`POST /api/v1/deposits` - Create deposit for the user](#create-deposit-for-the-user)
- [`GET /api/v1/deposits/{user-uuid}` - Get info about user's deposit](#get-user-deposit-info)
- [`POST /api/v1/deposits/{user-uuid}/transactions` - Make transaction to modify user's deposit funds](#make-user-deposit-transaction)

## Deposit data

```json
{
    "userUuid": "UUID",
    "deposit": "int",
    "creationDate": "timestamp"
}
```

## Transaction data

```json
{
    "amount": "int",
    "reason": "string",
    "partnerUuid"?: "UUID",
    "transactionDate": "timestamp"
}
```

##### Create deposit for the user
1. `POST /api/v1/deposits` - Create empty deposit for the specified user
    - Request Body:
        ```json
        {
            "userUuid": "13570e16-4d98-4823-b266-eeb3f4776eda"
        }
        ```
    - Responses:
        - `200` 
            ```json
            {
                "data": {
                    "userUuid": "13570e16-4d98-4823-b266-eeb3f4776eda",
                    "deposit": 0,
                    "creationDate": 1630753631
                }
            }
            ```
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "failed to convert provided userUuid field into UUID"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

##### Get user deposit info
2. `GET /api/v1/deposits/{user-uuid}` - Get user's deposit info
    - Request Body:
        (empty)
    - Responses:
        - `200` 
            ```json
            {
                "data": {
                    "userUuid": "13570e16-4d98-4823-b266-eeb3f4776eda",
                    "deposit": 0,
                    "creationDate": 1630753631
                }
            }
            ```
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "failed to convert provided userUuid field into UUID"
            }
            ```
        - `404`
            ```json
            {
                "error": "user not found",
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

##### Make user deposit transaction
3. `POST /api/v1/deposits/{user-uuid}/transactions` - Create new transaction on user's deposit to change funds
    - Request Body:
        ```json
        {
            "amount": 500,
            "reason": "add funds from credit card via application"
        }
        ```
        ```json
        {
            "amount": 500,
            "initiatorUserUuid": "13570e16-4d98-4823-b266-eeb3f4776eda",
            "reason": "user to user transaction"
        }
        ```
        ```json
        {
            "amount": -500,
            "reason": "system recommendation subscription"
        }
        ```
    - Responses:
        - `200` 
            ```json
            {
                "data": {
                    "amount": 500,
                    "reason": "add funds from credit card via application",
                    "transactionDate": 1630753631
                }
            }
            ```
            ```json
            {
                "data": {
                    "amount": 500,
                    "reason": "user to user transaction",
                    "partnerUuid": "13570e16-4d98-4823-b266-eeb3f4776eda",
                    "transactionDate": 1630753631
                }
            }
            ```
            ```json
            {
                "data": {
                    "amount": -500,
                    "reason": "system recommendation subscription",
                    "transactionDate": 1630753631
                }
            }
            ```
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "failed to convert provided userUuid field into UUID"
            }
            ```
            ```json
            {
                "error": "invalid data format",
                "description": "amount or reason field was not provided"
            }
            ```
            ```json
            {
                "error": "invalid request",
                "description": "initiator cannot withdraw funds from other user's deposit"
            }
            ```
        - `402`
            ```json
            {
                "error": "not enough funds",
                "description": "not enough funds on user's deposit to make withdraw with specified value"
            }
            ```
            ```json
            {
                "error": "not enough funds",
                "description": "not enough funds on initiator user's deposit to make funds transfer to another user"
            }
            ```
        - `404`
            ```json
            {
                "error": "user not found",
            }
            ```
            ```json
            {
                "error": "intiator user not found",
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```
