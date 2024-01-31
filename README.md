
# GO Authentication API

Authentication API using Golang, created with the aim of studying and practicing the language and frameworks.


## Features

- Login system using JWT and Cookies
- Protected Routes
- User CRUD
- Auto DB-Migration


## Stack used

**Back-end:** Golang, Go-fiber, Gorm, Golang-jwt and Postgres


## Roadmap

- PATCH route to edit user information

- Implement tests

- Implement front-end view using [HTMX](https://htmx.org/)

- 2FA

- Bruteforce login timeout

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

And you should use some client to test the routes, you can try [Postman](https://www.postman.com/downloads/)

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

#### Start application

```bash
  go run ./cmd/app/main.go
```

## Other commands

#### Stop docker compose
```bash
  docker-compose stop
```

#### Build application
```bash
  go build -o bin/go-authentication.exe ./cmd/app/main.go
```

#### Clean build files
```bash
  go clean
  rmdir /s /q bin
```
## API documentation

#### Login user

```http
  POST /login
```

| Parameter   | Type       | Description                           |
| :---------- | :--------- | :---------------------------------- |
| `email` | `string` | **Required**. User email |
| `password` | `string` | **Required**. User password |

Return a success message and set user token on browser Cookie

#### Logout user

```http
  POST /logout
```
Remove the token Cookie from browser, not letting user access to protected routes

#### Create user

```http
  POST /user
```

- Protected route, user should be logged-in


| Parameter   | Type       | Description                           |
| :---------- | :--------- | :---------------------------------- |
| `name` | `string` | **Required**. User name |
| `email` | `string` | **Required**. User email |
| `password` | `string` | **Required**. User password |

Return a success message and an object with the created user information

#### Get user by uuid

```http
  GET /user/uuid
```

- Protected route, user should be logged-in

#### Need to pass uuid on the query params eg:

```http
  GET localhost:8080/user/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee
```

Return a success message and an object with the user information

#### Get all users

```http
  GET /user
```

- Protected route, user should be logged-in


Return a success message and an object with all users information

#### Delete by uuid

```http
  DELETE /user/uuid
```

- Protected route, user should be logged-in

#### Need to pass uuid on the query params eg:

```http
  DELETE localhost:8080/user/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee
```

Return a success message and an object with the user information


## Authors

- [Artur Baldo Ramos](https://github.com/arturbaldoramos)

