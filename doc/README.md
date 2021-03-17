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
    "value": "a@a.fr",
    "active": true
  },
  {
    "key": "password",
    "value": "a441b15fe9a3cf56661190a0b93b9dec7d04127288cc87250967cf3b52894d11",
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
    "value": "a@b.fr",
    "active": true
  },
  {
    "key": "password",
    "value": "a441b15fe9a3cf56661190a0b93b9dec7d04127288cc87250967cf3b52894d11",
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
 - path: /users/a@a.fr
 - method: GET
 - passwordFieldType: password
 - bearerToken: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTE5NTAyMDUsIm5iZiI6MTYxNTk1MDIwNSwidXNlcmlkIjoiYUBhLmZyIn0.3ZF8QhgmoN_z6zF8r1PSjL2jTcJpdQUYzupsDtolcSM


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
 - path: /users/a@a.fr
 - method: PATCH
 - passwordFieldType: password
 - bearerToken: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTE5NTAyMDUsIm5iZiI6MTYxNTk1MDIwNSwidXNlcmlkIjoiYUBhLmZyIn0.3ZF8QhgmoN_z6zF8r1PSjL2jTcJpdQUYzupsDtolcSM


 - bodyParams:
```json
[
  {
    "key": "email",
    "value": "a@a.fr",
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
 - path: /users/a@a.fr
 - method: DELETE
 - passwordFieldType: password
 - bearerToken: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTE5NTAyMDUsIm5iZiI6MTYxNTk1MDIwNSwidXNlcmlkIjoiYUBhLmZyIn0.3ZF8QhgmoN_z6zF8r1PSjL2jTcJpdQUYzupsDtolcSM




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