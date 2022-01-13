# crypto-payments
Rest-API сервис для подключения платежей в крипто-валюте 

## Доступные API-методы

### /getNewInvoice

#### Запрос
* Метод: GET
* Тело запроса:
```
{
  "userId": int,
  "assert": string,
  "amount": int,
  "callbackUrl": string
}
```
#### Ответ
* Тело ответа:
```
{
  "invoiceAddress": string
}
```
