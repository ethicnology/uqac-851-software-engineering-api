
# prix-banque



## Indices

* [banks](#banks)

  * [create](#1-create)
  * [destroy](#2-destroy)
  * [index](#3-index)
  * [show](#4-show)
  * [update](#5-update)

* [invoices](#invoices)

  * [create](#1-create-1)
  * [destroy](#2-destroy-1)
  * [index](#3-index-1)
  * [show](#4-show-1)
  * [update](#5-update-1)

* [operations](#operations)

  * [index](#1-index)
  * [show](#2-show)

* [transfers](#transfers)

  * [create](#1-create-2)
  * [destroy](#2-destroy-2)
  * [index](#3-index-2)
  * [show](#4-show-2)
  * [update](#5-update-2)

* [users](#users)

  * [create](#1-create-3)
  * [destroy](#2-destroy-3)
  * [login](#3-login)
  * [show](#4-show-3)
  * [update](#5-update-3)
  * [verify](#6-verify)


--------


## banks



### 1. create



***Endpoint:***

```bash
Method: POST
Type: 
URL: https://dissidence.dev:9999/users/user1prixbanque@gmail.com/banks
```



### 2. destroy



***Endpoint:***

```bash
Method: DELETE
Type: RAW
URL: https://dissidence.dev:9999/users/user1prixbanque@gmail.com/banks/1
```



### 3. index



***Endpoint:***

```bash
Method: GET
Type: RAW
URL: https://dissidence.dev:9999/users/user1prixbanque@gmail.com/banks
```



### 4. show



***Endpoint:***

```bash
Method: GET
Type: RAW
URL: https://dissidence.dev:9999/users/user1prixbanque@gmail.com/banks/1
```



### 5. update



***Endpoint:***

```bash
Method: PATCH
Type: RAW
URL: https://dissidence.dev:9999/users/user1prixbanque@gmail.com/banks/1
```



***Body:***

```js        
{
    "balance": 5
}
```



## invoices



### 1. create



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: https://dissidence.dev:9999/users/user1prixbanque@gmail.com/banks/1/operations/invoices
```



***Body:***

```js        
{
    "amount": 11,
    "receiver_id": 1,
    "acquitted": "false",
    "due_date": "2021-05-11"
}
```



### 2. destroy



***Endpoint:***

```bash
Method: DELETE
Type: 
URL: https://dissidence.dev:9999/users/user1prixbanque@gmail.com/banks/1/invoices/1
```



### 3. index



***Endpoint:***

```bash
Method: GET
Type: 
URL: https://dissidence.dev:9999/users/user1prixbanque@gmail.com/banks/1/invoices
```



### 4. show



***Endpoint:***

```bash
Method: GET
Type: 
URL: https://dissidence.dev:9999/users/user1prixbanque@gmail.com/banks/1/invoices/1
```



### 5. update



***Endpoint:***

```bash
Method: PATCH
Type: RAW
URL: https://dissidence.dev:9999/users/user1prixbanque@gmail.com/banks/1/invoices/1
```



***Body:***

```js        
{
    "amount": 17,
    "acquitted": "false",
    "due_date": "2021-05-11"
}
```



## operations



### 1. index



***Endpoint:***

```bash
Method: GET
Type: 
URL: https://dissidence.dev:9999/users/user1prixbanque@gmail.com/banks/1/operations
```



### 2. show



***Endpoint:***

```bash
Method: GET
Type: 
URL: https://dissidence.dev:9999/users/user1prixbanque@gmail.com/banks/1/operations/1
```



## transfers



### 1. create



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: https://dissidence.dev:9999/users/user1prixbanque@gmail.com/banks/1/transfers
```



***Body:***

```js        
{
    "amount": 11,
    "receiver_id": 1,
    "scheduled": "false",
    "date": "2006-01-02",
    "question": "To be or not ?",
    "answer": "To be..."
}
```



### 2. destroy



***Endpoint:***

```bash
Method: DELETE
Type: 
URL: https://dissidence.dev:9999/users/user1prixbanque@gmail.com/banks/1/transfers/1
```



### 3. index



***Endpoint:***

```bash
Method: GET
Type: 
URL: https://dissidence.dev:9999/users/user1prixbanque@gmail.com/banks/1/transfers
```



### 4. show



***Endpoint:***

```bash
Method: GET
Type: 
URL: https://dissidence.dev:9999/users/user1prixbanque@gmail.com/banks/1/transfers/11
```



### 5. update



***Endpoint:***

```bash
Method: PATCH
Type: RAW
URL: https://dissidence.dev:9999/users/user1prixbanque@gmail.com/banks/1/transfers/1
```



***Body:***

```js        
{
    "amount": 20,
    "scheduled": "true",
    "date": "2021-05-11",
    "question": "To be or not ?",
    "answer": "To be..."
}
```



## users



### 1. create



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: https://dissidence.dev:9999/auth/register
```



***Body:***

```js        
{
    "email":"{{email}}",
    "password": "{{password}}",
    "first_name": "User1",
    "last_name": "Lambda"
}
```



### 2. destroy



***Endpoint:***

```bash
Method: DELETE
Type: RAW
URL: https://dissidence.dev:9999/users/user1prixbanque@gmail.com
```



### 3. login



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: https://dissidence.dev:9999/auth/login
```



***Body:***

```js        
{
    "email":"{{email}}",
    "password": "{{password}}"
}
```



### 4. show



***Endpoint:***

```bash
Method: GET
Type: RAW
URL: https://dissidence.dev:9999/users/user1prixbanque@gmail.com
```



### 5. update



***Endpoint:***

```bash
Method: PATCH
Type: RAW
URL: https://dissidence.dev:9999/users/user1prixbanque@gmail.com
```



***Body:***

```js        
{
    "email":"{{email}}",
    "first_name": "Léon",
    "last_name": "Debally"
}
```



### 6. verify



***Endpoint:***

```bash
Method: GET
Type: RAW
URL: https://dissidence.dev:9999/users/user1prixbanque@gmail.com/verify/GRVEmbM2hVAIb3VJKyNa-AOzcZpR3xB-
```



***Available Variables:***

| Key | Value | Type |
| --- | ------|-------------|
| URL | https://dissidence.dev:9999 |  |
| email | user1prixbanque@gmail.com |  |
| password | 1c002df4bdb6cd88651c6c2e92a1d63ae0c016b2f96b6ffd452c0a6219dd33b7 |  |
| bank_id | 1 |  |
| operation_id | 1 |  |



---
[Back to top](#prix-banque)
> Made with &#9829; by [thedevsaddam](https://github.com/thedevsaddam) | Generated at: 2021-04-22 16:36:36 by [docgen](https://github.com/thedevsaddam/docgen)
