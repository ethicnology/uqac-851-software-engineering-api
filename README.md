## About The Project
![](https://github.com/ethicnology/uqac-851-software-engineering-api/blob/main/doc/logo.png "Screenshot")

This is a school project.  
All the specifications are specified in [doc/Projet_Pratique.pdf](https://github.com/ethicnology/uqac-851-software-engineering-api/blob/main/doc/Projet_Pratique.pdf)

* [Docker](https://www.docker.com)
* [Golang](https://golang.org)  
* [Goyave](https://goyave.dev) 
* [MariaDB](https://mariadb.org) 

### :open_file_folder: Directory Structure
```
.
├── database
│   ├── model                // ORM models and Generators for database testing
│       └── ...

├── http
│   ├── controller           // Business logic of the application
│   │   └── ...
│   ├── middleware           // Logic executed before or after controllers
│   │   └── ...
│   └── route
│       └── route.go         // Routes definition
│
├── test                     // Unit and Functional tests
|   └── func_test.go
|   └── unit_test.go
|
├── .gitignore
├── config.example.json      // Example config for local development
├── config.test.json         // Config file used for tests
├── docker-compose.yml       // Build local architecture file used for tests
├── docker-compose.test.yml  // Build architecture to execute tests
├── Dockerfile               // API Dockerfile
├── go.mod
└── main.go                  // Application entrypoint
```


## :rocket: Getting Started
To get a local copy up and running follow these simple example steps.
### :page_with_curl: Prerequisites
All the project architecture is dockerized : [Docker](https://www.docker.com/products/docker-desktop).  
You only need docker installed in your OS in order to host all the components without any extra-installation.  

### :construction_worker: Installation
#### Clone
```sh
git clone https://github.com/ethicnology/uqac-851-software-engineering-api.git
```
#### Configuration
Rename **config.example.json** as **config.json** and change values if needed.  
All the fields are documented by [Goyave Framework](https://goyave.dev/guide/configuration.html)
```sh
mv config.example.json config.json
```

Rename **.env.example** as **.env** and specify **the password for smtp mail** (ask to students) :
```sh
mv .env.example .env
```

```ini
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
MAIL_USER=prixbanque@gmail.com
MAIL_PASS=<AskToStudents>
```


### :whale: Run with docker
```sh
docker-compose up
# Available on 172.20.0.3:1984
```

### :pray: Tests
```sh
docker-compose -f docker-compose.test.yml up --abort-on-container-exit --remove-orphans
# Exit when tests are finished
```

### :runner: Usage
#### Hoppscotch.io
To play with your API, import the collection in /doc in [Hoppscotch](https://hoppscotch.io/).  
All available endpoints are documented with examples.  

#### cURL
Instead you can try it with your CLI using **cURL**.  

##### Windows
Windows force you to mask doublequotes with backslash like this : 
```powershell
curl -d "{\"email\":\"sensei@uqac.ca\",\"password\":\"Str0ng\"}" -H "Content-Type: application/json" -X POST http://172.20.0.3:1984/auth/register
```
##### Linux
```sh
curl -d '{"email":"sensei@uqac.ca","password":"Str0ng"}' -H 'Content-Type: application/json' -X POST http://172.20.0.3:1984/auth/register
# {"id":9}
```

```sh
curl -d '{"email":"sensei@uqac.ca","password":"Str0ng"}' -H 'Content-Type: application/json' -X POST http://172.20.0.3:1984/auth/login
# {"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTM0NzkyNDQsIm5iZiI6MTYxNzQ3OTI0NCwidXNlcmlkIjoic2Vuc2VpQHVxYWMuY2EifQ.aMRWeebCTfJyUPfsUz5H8Ng1x1L1T10hSKpXoVdyPUY"}
```

```sh
TOKEN=$YourToken;curl -H 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" -X GET http://172.20.0.3:1984/users/sensei@uqac.ca
# {"created_at":"2021-04-03T19:47:19.718Z","updated_at":"2021-04-03T19:47:19.718Z","deleted_at":null,"id":10,"email":"sensei@uqac.ca","first_name":"","last_name":""}
```

### :hammer: Build from scratch
#### Prerequisites
If you want to build the project from scratch without docker you need to install few tools.
* Go 1.13+
* mariadb-server

Once you've installed Go and MariaDB.
You need to create a MariaDB user and a database:
```sql
mysql> CREATE USER 'goyave'@'localhost' IDENTIFIED BY 'secret';
mysql> CREATE DATABASE goyave;
mysql> GRANT ALL PRIVILEGES ON 'goyave'.* TO 'goyave'@'localhost';
mysql> FLUSH PRIVILEGES;
```
You can change theses values but you will need to update **config.json**.

#### Build
From root directory, execute :
```sh
go build
```
It will output an executable which have the same name as the project, make it executable :
```sh
chmod +X uqac-851-software-engineering-api
```
Then, run :
```sh
./uqac-851-software-engineering-api
```
## :mag: Contact
ethicnology - [@ethicnology](https://twitter.com/ethicnology)  
Project Link: [https://github.com/851-software-engineering](https://github.com/851-software-engineering)
