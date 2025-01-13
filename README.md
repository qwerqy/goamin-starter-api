# Go API server starter

## Structure

`bin` is where compiled code will be.

`cmd` is where the executable will be. Main entry point, which is the `api` and
`migrate`. `api` will have everything server related. `migrate` is where all
migration related code will live.

`docs` is where swagger documentation will live.

`internal` is where internal code will be stored. These will not be
visible/exportable out of the project scope and will only be used within the
project scope.

`scripts` is where setting up the server script will be.

### Additional

`web` that contains frontend code. (React, Svelte, NextJS, etc)

## Principles

### Separation of concerns

Each level in your program should be separated by a clear barrier, the transport
layer, the service layer, the storage layer...

### Dependency Inversion Principle (DIP)

You're injecting the dependencies in your layers. You don't directly call them.
Why? It promotes loose coupling and makes it easier to test your programs.

### Adaptability to Change

By organizing your code in a modular and flexible way, you can more easily
introduce new features, refactor existing code, and respond to evolving business
requirements.

Your systems should be easy to change, if you have to change alot of existing
code to add a new feature you're doing it wrong.

### Focus on Business Value

And finally, focus on delivering value to your users, they are the ones who will
be paying your bills at the end of the month.. So focus on business value.

## Replace Placeholders

- Replace package placeholder in `go.mod` with your own package name.
- Replace `github.com/qwerqy/api-starter` with your own package name.
- Replace `api-starter` with your own project name.

## Install dependencies

```bash
go mod tidy
```

## Development

```bash
air
```

### Environment variables

You can set environment variables in `.envrc` file. Refer to `.env.example` for
more information.

I use [direnv](https://direnv.net/) to manage environment variables.

```bash
direnv allow .
```

### Database & Caching

This project is built with the adapter pattern in mind. My preferred database is
PostgreSQL, but you can easily swap it out for any other database that has a
driver for it. You can find the database adapter in `internal/db/db.go`.

For caching I use Redis, but you can easily swap it out for any other caching
solution that has a driver for it. You can find the cache adapter in
`internal/cache/cache.go`.

For migrations I use
[golang-migrate](https://github.com/golang-migrate/migrate). Refer to the
Makefile for more information.

## Mailer

This project is built with the adapter pattern in mind. My preferred mailer is
SendGrid, but you can easily swap it out for any other mailer that has a driver
for it. You can find the mailer adapter in `internal/mailer/mailer.go`.

## Authentication

My preferred authentication solution is JWT, but you can easily swap it out for
any other authentication solution that has a driver for it. You can find the
authentication adapter in `internal/auth/jwt.go`.

## Generate docs

Generate docs run automatically when you run `air`.

```bash
make gen-docs
```
