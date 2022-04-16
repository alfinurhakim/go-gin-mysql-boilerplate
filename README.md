# go-gin-mysql-boilerplate
Golang Restful CRUD & Authentication API using Gin Framework, Gorm, and MySQL.

This boilerplate has been equipped with :
- Validation
- Pagination
- JWT Authentication
- Documentation Swagger
- Handlers CORS

### Structure

```
.
├── config
│   └── development.yaml
│   └── production.yaml
│   └── stage.yaml
│   └── test.yaml
├── controllers
│   └── auth_controller.go
│   └── books_controller.go
│   └── default_controller.go
│   └── users_controller.go
├── docs
│   └── swagger.json
│   └── swagger.yaml
├── helper
│   ├── helper_responses.go
│   ├── paginations.go
├── lib
│   ├── config
│   │   ├── config.go
├── middlewares
│   ├── auth.go
├── migrations
│   ├── migrations.go
├── models
│   ├── books.go
│   ├── users.go
├── repositories
│   ├── auth_repository.go
│   ├── books_repository.go
│   ├── users_repository.go
├── responses
│   ├── responses.go
├── router
│   ├── router.go
│   ├── server.go
├── templates
│   ├── index.tmpl
├── test
├── utils
│   ├── token.go
├── validations
│   ├── auth_validation.go
│   ├── books_validation.go
├── .env
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── README.md
└── main.go

```

## Project Setup
### Set your .env in (env-example)
```
DB_HOST = localhost
DB_PORT = 8889
DB_USER = root
DB_PASSWORD = root
DB_NAME = db-golang-crud
MODE = development

TOKEN_HOUR_LIFESPAN=24
API_SECRET=ZGFsYW0gc3VyYXQgYWwtTWFpZGFoIGF5YXQgMzg6IExha2ktbGFraSB5YW5nIG1lbmN1cmkgZGFuIHBlcmVtcHVhbiB5YW5nIG1lbmN1cmksIHBvdG9uZ2xhaCB0YW5nYW4ga2VkdWFueWEgKHNlYmFnYWkpIHBlbWJhbGFzYW4gYmFnaSBhcGEgeWFuZyBtZXJla2Ega2VyamFrYW4gZGFuIHNlYmFnYWkgc2lrc2FhbiBkYXJpIEFsbGFoLg==

; For Migration :
; -> give CREATE_MIGRATION_ALL = 1 to migration all table
; -> give CREATE_MIGRATION_ALL = 2 to selection table migration & set 1 to selecting table 

CREATE_MIGRATION_ALL = 1
CREATE_MIGRATION_USERS = 0
CREATE_MIGRATION_BOOKS = 0
```
### Dependencies
- gin v1.5.0
- cors v0.0.0-20170318125340-cf4846e6a636
- gorm v1.23.4
- mysql v1.3.2
- crypto v0.0.0-20210921155107-089bfa567519
- swaggo/gin-swagger v1.2.0
- swaggo/files v0.0.0-20190704085106-630677cd5c14

### How To Run
```
go run main.go
```
## Results

### http://localhost:8000
<img width="1437" alt="Screen Shot 2022-04-16 at 22 02 52" src="https://user-images.githubusercontent.com/48195224/163680061-ef00d7ae-7038-4432-81d5-173e220481d5.png">

### http://localhost:8000/api-docs/index.html
![screencapture-localhost-8000-api-docs-index-html-2022-04-16-21_42_50](https://user-images.githubusercontent.com/48195224/163680220-92707358-6e85-4a52-9756-488f3c62da9f.png)




