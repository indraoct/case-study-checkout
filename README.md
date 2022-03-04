# case-study-checkout
A case study about module checkout in online shop

## Content

- [About Application](#about-application)
- [Technology Stack](#technology-stack)
- [Commands](#commands)
- [Project Structure](#project-structure)
- [Healthchecks](#healthchecks)
- [Unit Test](#unittest)

## About application
This application build using Golang. I used SOLID for the architecture. This application is API for function :
1. Add to cart
2. Checkout

## Technology Stack

Following is technology stack that is used in this service

| Name | Version | Description |
|------|---------|-------------|
| [Golang](https://golang.org/) | >= 1.14 | Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. |
| [Docker](https://www.docker.com/) | any | Enterprise Container Platform for High-Velocity Innovation. |


## Commands


First, install golang dependency by using go module:

```bash
go mod tidy
```

then, run the app by using this command:

```bash
go run main.go
```

webserver is going to running on `localhost:9000`.

## Project Structure
| Folder Name | Contain | Description |
|------|---------|-------------|
| constants |constanta| define static variable like : "PROMO_TYPE" |
|  | any | Enterprise Container Platform for High-Velocity Innovation. |


## Healthchecks

This service exposes a basic healthcheck at `/ping`:

```
GET /checkout/ping HTTP/1.1
host: localhost:9000

HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: 5

PONG!
```