## About The Project
![](https://github.com/ethicnology/uqac-851-software-engineering-api/blob/main/doc/logo.png "Screenshot")

This is a school project.  
All the specifications are specified in [doc/Projet_Pratique.pdf](https://github.com/ethicnology/uqac-851-software-engineering-api/blob/main/doc/Projet_Pratique.pdf)

### Built With
* [Golang](https://golang.org)  
* [Goyave](https://goyave.dev)  
Directory Structure
```
.
├── database
│   ├── model                // ORM models
│   |   └── ...
│   └── seeder               // Generators for database testing
│       └── ...
├── http
│   ├── controller           // Business logic of the application
│   │   └── ...
│   ├── middleware           // Logic executed before or after controllers
│   │   └── ...
│   ├── validation
│   │   └── validation.go    // Custom validation rules
│   └── route
│       └── route.go         // Routes definition
│
├── test                     // Functional tests
|   └── ...
|
├── .gitignore
├── .golangci.yml            // Settings for the Golangci-lint linter
├── config.example.json      // Example config for local development
├── config.test.json         // Config file used for tests
├── go.mod
└── main.go                  // Application entrypoint
```


<!-- GETTING STARTED -->
## Getting Started
To get a local copy up and running follow these simple example steps.

### Prerequisites
* Golang
  ```sh
    Go 1.13+
    Go modules
  ```

### Installation
1. Clone the repo
   ```sh
   git clone https://github.com/ethicnology/uqac-851-software-engineering-api.git
   ```
2. Create your own `config.json`
   ```JSON
{
    "app": {
        "name": "PrixBanque-api",
        "environment": "localhost",
        "debug": true,
        "defaultLanguage": "en-US"
    },
    "server": {
        "host": "127.0.0.1",
        "maintenance": false,
        "protocol": "http",
        "domain": "",
        "port": 8080,
        "httpsPort": 8081,
        "timeout": 10,
        "maxUploadSize": 10,
        "tls": {
            "cert": "",
            "key": ""
        }
    },
    "database": {
        "connection": "mysql",
        "host": "127.0.0.1",
        "port": 3306,
        "name": "PrixBanque-db",
        "username": "root",
        "password": "root",
        "options": "charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true&loc=Local",
        "maxOpenConnections": 20,
        "maxIdleConnections": 20,
        "maxLifetime": 300,
        "autoMigrate": false
    },
    "auth": {
        "jwt": {
            "expiry": 3600,
            "secret": "secret"
        }
    }    
}
   ```
3. Install NPM packages
   ```sh
   go run .
   ```

## Contact
ethicnology - [@ethicnology](https://twitter.com/ethicnology)  
Project Link: [https://github.com/851-software-engineering](https://github.com/851-software-engineering)
