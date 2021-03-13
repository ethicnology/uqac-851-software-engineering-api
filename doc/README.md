# API Documentation

## prix-banque
### Register
 - url: http://127.0.0.1:1984
 - path: /auth/register
 - method: POST
 - passwordFieldType: password


 - bodyParams:
```json
[
  {
    "key": "email",
    "value": "b@a.fr",
    "active": true
  },
  {
    "key": "password",
    "value": "rand0m",
    "active": true
  }
]
```

 - contentType: application/json
 - requestType: curl
 - preRequestScript:
```javascript
// pw.env.set('variable', 'value');
```
 - testScript:
```javascript
// pw.expect('variable').toBe('value');
```
---
### Login (JWT)
 - url: http://127.0.0.1:1984
 - path: /auth/login
 - method: POST
 - passwordFieldType: password


 - bodyParams:
```json
[
  {
    "key": "email",
    "value": "b@a.fr",
    "active": true
  },
  {
    "key": "password",
    "value": "rand0m",
    "active": true
  }
]
```

 - contentType: application/json
 - requestType: curl
 - preRequestScript:
```javascript
// pw.env.set('variable', 'value');
```
 - testScript:
```javascript
// pw.expect('variable').toBe('value');
```
---
### Get user
 - url: http://127.0.0.1:1984
 - path: /users/3
 - method: GET
 - passwordFieldType: password
 - bearerToken: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTU2Nzg5NDksIm5iZiI6MTYxNTY3NTM0OSwidXNlcmlkIjoiYkBhLmZyIn0.HVkWBdpL1lvvFuHHlMoq_xy2jD5VMslCo8sG_YRKVg0


 - bodyParams:
```json
[
  {
    "key": "email",
    "value": "b@a.fr",
    "active": true
  },
  {
    "key": "password",
    "value": "rand0m",
    "active": true
  }
]
```
 - requestType: curl
 - preRequestScript:
```javascript
// pw.env.set('variable', 'value');
```
 - testScript:
```javascript
// pw.expect('variable').toBe('value');
```
---
### Get users
 - url: http://127.0.0.1:1984
 - path: /users
 - method: GET
 - passwordFieldType: password
 - bearerToken: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTU2NzgzODUsIm5iZiI6MTYxNTY3NDc4NSwidXNlcmlkIjoiYUBhLmZyIn0.pC2FzsRiV8KTTZy65KZD-TplWyBlb9FnpGiSaQfbyhU


 - bodyParams:
```json
[
  {
    "key": "email",
    "value": "b@a.fr",
    "active": true
  },
  {
    "key": "password",
    "value": "rand0m",
    "active": true
  }
]
```
 - requestType: curl
 - preRequestScript:
```javascript
// pw.env.set('variable', 'value');
```
 - testScript:
```javascript
// pw.expect('variable').toBe('value');
```
---
### Patch user
 - url: http://127.0.0.1:1984
 - path: /users/1
 - method: PATCH
 - passwordFieldType: password
 - bearerToken: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTU2NzgzODUsIm5iZiI6MTYxNTY3NDc4NSwidXNlcmlkIjoiYUBhLmZyIn0.pC2FzsRiV8KTTZy65KZD-TplWyBlb9FnpGiSaQfbyhU


 - bodyParams:
```json
[
  {
    "key": "email",
    "value": "b@a.fr",
    "active": true
  },
  {
    "key": "first_name",
    "value": "Léon",
    "active": true
  },
  {
    "key": "last_name",
    "value": "De la rue",
    "active": true
  }
]
```

 - contentType: application/json
 - requestType: curl
 - preRequestScript:
```javascript
// pw.env.set('variable', 'value');
```
 - testScript:
```javascript
// pw.expect('variable').toBe('value');
```
---
### Delete user
 - url: http://127.0.0.1:1984
 - path: /users/1
 - method: DELETE
 - passwordFieldType: password
 - bearerToken: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTU2NzgzODUsIm5iZiI6MTYxNTY3NDc4NSwidXNlcmlkIjoiYUBhLmZyIn0.pC2FzsRiV8KTTZy65KZD-TplWyBlb9FnpGiSaQfbyhU




 - contentType: application/json
 - requestType: curl
 - preRequestScript:
```javascript
// pw.env.set('variable', 'value');
```
 - testScript:
```javascript
// pw.expect('variable').toBe('value');
```
---

<br/>
