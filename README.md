
# GO Authentication API

Authentication API using Golang, created with the aim of studying and practicing the language and frameworks.

## Peek
![Peek](/pkg/utils/peek.png "Peek")

## Features

- SSR with templates and htmx
- Login system using JWT and Cookies
- Protected Routes
- User CRUD
- Auto DB-Migration


## Stack used

**Back-end:** Golang, Go-fiber, Gorm.

**Front-end:** Templ, HTMX, Flowbite and Tailwind.

**Database:** Postgres


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
Firstly you need to make sure that Golang, Docker and NPM are installed on your machine,
if you don't have them, follow the link to their installation:

[Golang](https://go.dev/doc/install)

[Docker](https://www.docker.com/)

[NPM](https://www.npmjs.com/)


----
#### Clone the project

```bash
  git clone https://github.com/arturbaldoramos/go-authentication.git
```

#### Enter project directory

```bash
  cd go-authentication
```

#### Install dependencies

```bash
  go mod tidy
  npm install
```

#### Generate css styles

```bash
  npx tailwindcss -o ./pkg/static/output.css
```


#### Run docker compose

```bash
  docker-compose up -d
```


## Useful information

#### Hot-reload to develop?
Install [AIR](https://github.com/cosmtrek/air) - config file already on the project

Run command:
```bash
  air
```

---

#### Want to create new templates and use them?
Install [Templ](https://templ.guide/) - model files in pkg/template

Run command to build template
```bash
  templ generate
```

---
#### Want to update htmx version?
Go to their [Website](https://htmx.org/docs/#installing) and look to download a copy. Replace the file inside `/pkg/static/htmx.min.js`, please keep the same filename. 

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

