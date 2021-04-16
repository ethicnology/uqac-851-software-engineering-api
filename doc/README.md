
# prix-banque



## Indices

* [banks](#banks)

  * [Create](#1-create)
  * [Destroy](#2-destroy)
  * [Index](#3-index)
  * [Show](#4-show)
  * [Update](#5-update)

* [operations](#operations)

  * [Create Invoice](#1-create-invoice)
  * [Create Transfer](#2-create-transfer)
  * [Index Invoices](#3-index-invoices)
  * [Index Operations](#4-index-operations)
  * [Index Transfers](#5-index-transfers)

* [users](#users)

  * [Create](#1-create-1)
  * [Destroy](#2-destroy-1)
  * [Login](#3-login)
  * [Show](#4-show-1)
  * [Update](#5-update-1)


--------


## banks



### 1. Create



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: http://172.20.0.3:1984/users/a@a.fr/banks
```



***Body:***

```js        
{
    "balance": 100
}
```



### 2. Destroy



***Endpoint:***

```bash
Method: DELETE
Type: RAW
URL: http://172.20.0.3:1984/users/a@a.fr/banks/4
```



### 3. Index



***Endpoint:***

```bash
Method: GET
Type: RAW
URL: http://172.20.0.3:1984/users/a@a.fr/banks
```



### 4. Show



***Endpoint:***

```bash
Method: GET
Type: RAW
URL: http://172.20.0.3:1984/users/a@a.fr/banks/1
```



### 5. Update



***Endpoint:***

```bash
Method: PATCH
Type: RAW
URL: http://172.20.0.3:1984/users/a@a.fr/banks/1
```



***Body:***

```js        
{
    "balance": 5
}
```



## operations



### 1. Create Invoice



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: http://172.20.0.3:1984/users/a@a.fr/banks/1/operations/invoices
```



***Body:***

```js        
{
    "amount": 11,
    "receiver_id": 2,
    "acquitted": "false",
    "due_date": "2021-05-11"
}
```



### 2. Create Transfer



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: http://172.20.0.3:1984/users/a@a.fr/banks/1/operations/transfers
```



***Body:***

```js        
{
    "amount": 11,
    "receiver_id": 1,
    "scheduled": "false",
    "date": "2021-05-11",
    "question": "To be or not ?",
    "answer": "To be..."
}
```



### 3. Index Invoices



***Endpoint:***

```bash
Method: GET
Type: 
URL: http://172.20.0.3:1984/users/a@a.fr/banks/1/operations/invoices
```



### 4. Index Operations



***Endpoint:***

```bash
Method: GET
Type: 
URL: http://172.20.0.3:1984/users/a@a.fr/banks/1/operations
```



### 5. Index Transfers



***Endpoint:***

```bash
Method: GET
Type: 
URL: http://172.20.0.3:1984/users/a@a.fr/banks/1/operations/transfers
```



## users



### 1. Create



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: http://172.20.0.3:1984/auth/register
```



***Body:***

```js        
{
    "email":"{{email}}",
    "password": "{{password}}"
}
```



### 2. Destroy



***Endpoint:***

```bash
Method: DELETE
Type: RAW
URL: http://172.20.0.3:1984/users/a@a.fr
```



### 3. Login



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: http://172.20.0.3:1984/auth/login
```



***Body:***

```js        
{
    "email":"{{email}}",
    "password": "{{password}}"
}
```



### 4. Show



***Endpoint:***

```bash
Method: GET
Type: RAW
URL: http://172.20.0.3:1984/users/a@a.fr
```



### 5. Update



***Endpoint:***

```bash
Method: PATCH
Type: RAW
URL: http://172.20.0.3:1984/users/a@a.fr
```



***Body:***

```js        
{
    "email":"{{email}}",
    "first_name": "Léon",
    "last_name": "Debally"
}
```



***Available Variables:***

| Key | Value | Type |
| --- | ------|-------------|
| URL | http://172.20.0.3:1984 |  |
| email | a@a.fr |  |
| password | a441b15fe9a3cf56661190a0b93b9dec7d04127288cc87250967cf3b52894d11 |  |



---
[Back to top](#prix-banque)
> Made with &#9829; by [thedevsaddam](https://github.com/thedevsaddam) | Generated at: 2021-04-16 00:01:33 by [docgen](https://github.com/thedevsaddam/docgen)
