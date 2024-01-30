
# GO Authentication API

Authentication API using Golang, created with the aim of studying and practicing the language and frameworks.


## Features

- Login system using JWT and Cookies
- Protected Routes
- User CRUD
- Auto DB-Migration


## Stack used

**Back-end:** Golang, Go-fiber, Gorm, Golang-jwt and Postgres


## Environment variables

To run this project, you will need to rename the .env.example file to .env, or create your own with the following variables.

`API_PORT = :8080`                       Port to serve api

`JWT_SECRET = secret`                    Jwt secret to sign token (generate a secure one)

`JWT_EXPIRATION = 24`                    Jwt token expiration time (in hours)



`DB_HOST = localhost`                    Database host

`DB_PORT = 5432`                         Database port

`DB_USER = postgres`                     Database username

`DB_PASSWORD = password`                 Database password

`DB_NAME = database_name`                Database name


## Running locally

#### Required
Firstly you need to make sure that Golang and Docker are installed on your machine,
if you don't have them, follow the link to their installation:

[Golang](https://go.dev/doc/install)

[Docker](https://www.docker.com/)

#### Optional

You can download   [Make](https://gnuwin32.sourceforge.net/packages/make.htm) to         easily   run project commands

----
#### Clone the project

```bash
  git clone https://github.com/arturbaldoramos/go-authentication.git
```

#### Enter project directory

```bash
  cd my-project
```

#### Install dependencies

```bash
  go mod tidy
```

#### Run docker compose

```bash
  docker-compose up -d
```

#### Start aplication

```bash
  go run ./cmd/app/main.go
```

## Other commands

#### Stop docker compose
```bash
  docker-compose stop
```

#### Build aplication
```bash
  go build -o bin/go-authentication.exe ./cmd/app/main.go
```

#### Clearn build files
```bash
  go clean
  rmdir /s /q bin
```
## Autores

- [Artur Baldo Ramos](https://github.com/arturbaldoramos)

