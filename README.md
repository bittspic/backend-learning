# Simple bank

A simple bank application

## Built with

- Language: Go
- Framework: Gin
- Database: Postgres
- Virtualization: Docker & Docker Compose
- Design DB schema and generate SQL code: dbdiagram.io
- DB migration: golang-migrate
- Generate CRUD Golang code from SQL: sqlc
- CI: Github Actions
- Load config: viper
- Hash password: Bcrypt
- Authenticaion: JWT, PASETO

## Feature

- Login
- Create user
- Create and manage accounts's information
- Record balance changes to each of accounts
- Perform money transfer between two accounts in the transaction

## Setup local development

- Go
- Docker
- golang-migrate
- sqlc
- golang mock

## Run

- Start all service: `docker compose up`
- Start postgres: `make postgres`
- Start Go server: `make server`
- Run unit test: `make test`
- Create DB: `make createdb`
- Drop DB: `make dropdb`
- Migrate up: `make migrateup`
- Migrate up one version: `make migrateup1`
- Migrate down: `make migratedown`
- Migrate down one version: `make migratedown1`
- Generate CRUD code: `make sqlc`
- Generate mock interface: `make mock`