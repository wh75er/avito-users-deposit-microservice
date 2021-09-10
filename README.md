# Тестовое задание на позицию стажера-бекендера

## Микросервис для работы с балансом пользователей

Здесь представлена реализация, документация, api микросервиса для работы с балансом пользователя

Оглавление:

- [Документация API](#документация-api)
- [Запуск](#запуск)
- [Тестовые запросы для "потыкать"(curl)](#тестовые-запросы)

## Документация API

Документация api изложена на английском и представлена в формате MD [здесь](docs/api.md).

## Таблицы БД

В качестве базы данных используется Postgresql

Здесь представлены таблицы БД, которые используются в проекте:


```sql
CREATE TABLE IF NOT EXISTS Deposits(
    Id serial PRIMARY KEY,
    UserUuid UUID NOT NULL UNIQUE,
    Deposit INT NOT NULL,
    CreationDate TIMESTAMP NOT NULL
);
```

```sql
CREATE TABLE IF NOT EXISTS Transactions(
    Id serial PRIMARY KEY,
    DepositId INT REFERENCES Deposits(Id),
    OwnerUuid UUID NOT NULL,
    Amount INT NOT NULL,
    Reason VARCHAR(250) NOT NULL,
    PartnerUuid UUID,
    TransactionDate TIMESTAMP NOT NULL
);
```

## Запуск

Для запуска:

```
sudo docker-compose up
```

## Тестовые запросы

Зачислить средства пользователю по указанному UUID(Успех):

```
curl -v -H "Content-Type: application/json" -X POST localhost:3000/api/v1/deposits/c4a07944-a0b3-4c10-9803-6af93971b3d5/transactions -d '{"amount": 400, "reason": "credit card via online bank"}'
```

Посмотреть информацию о счете для указанного пользователя(Успех):

```
curl -v -X GET localhost:3000/api/v1/deposits/c4a07944-a0b3-4c10-9803-6af93971b3d5
```

Попытаться сделать перевод денег от пользователя другому пользователю(инициатор транзакции не имеет счета):

```
curl -v -H "Content-Type: application/json" -X POST localhost:3000/api/v1/deposits/c4a07944-a0b3-4c10-9803-6af93971b3d5/transactions -d '{"amount": 5000, "reason": "user to user transaction", "initiatorUserUuid": "18866f0a-4a34-46a0-aff0-69d3334f3c23"}'
```

Попытаться закинуть деньги с несуществующего счета на существующий(Не успех):

```
curl -v -H "Content-Type: application/json" -X POST localhost:3000/api/v1/deposits/18866f0a-4a34-46a0-aff0-69d3334f3c23/transactions -d '{"amount": 2000, "reason": "user to  user transaction", "initiatorUserUuid": "18866f0a-4a34-46a0-aff0-69d3334f3c23"}'
```

Создать счет для другого пользователя(Успех):

```
curl -v -H "Content-Type: application/json" -X POST localhost:3000/api/v1/deposits/18866f0a-4a34-46a0-aff0-69d3334f3c23/transactions -d '{"amount": 2000, "reason": "bank money transfer"}'
```

Вывести деньги со счета указанного пользователя(Успех):

```
curl -v -H "Content-Type: application/json" -X POST localhost:3000/api/v1/deposits/18866f0a-4a34-46a0-aff0-69d3334f3c23/transactions -d '{"amount": -500, "reason": "avito account+ month subscription"}'
```

### Проблемы при проектировании api и других различных фич:

2. Изменение баланса пользователя, которое тоже можно сделать разными методами:
    - PATCH запрос на изменение депозита для конкретного пользователя
    - Использование отдельных ручек в api для добавления средств и списания
      средств, например, deposit/{user-uuid}/raise и
      deposit/{user-uuid}/withdraw
    - Использование одной api ручки POST /deposit/{user-uuid}/change, где будет
      приходить int значение с - или + значением суммы(если <0, то списание,
      если >0, то зачисление)
    - В конце пришел к тому, что лучше сделать отдельную сущность для транзакций.
      Таким образом, при изменении счета следует просто добавлять новую запись
      в транзакции. Минусом такого подхода - толстый обработчик добавления транзакции
3. К чему должны быть привязаны транзакции в БД?
    - В идеале таблица транзакций должна ссылаться на какой-то счет(по идентификатору), но
      так как в данной архитектуре у пользователя может быть только один счет, можно сделать
      привязку к идентификатору юзера. Так как открытие нескольких счетов довольно специфическая
      фича, которая используется в основном только в банковских приложениях, я привяжу транзакцию
      к идентификаторую юзера, чей счет в данный момент редактируется и является единственным и 
      уникальным
4. Для обеспечения синхронизации между транзакциями к разным репозиториям нужно выделить интерфейс
    транзактора и передавать его из конкретного репозитория в юзкейс. Тогда любой репозиторий
    мог бы реализовывать его транзакции с возможностью коммитов и роллбеков(не успел это сделать)
