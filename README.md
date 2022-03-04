# case-study-checkout
A case study about module checkout in online shop

## Content

- [About Application](#about-application)
- [Technology Stack](#technology-stack)
- [Commands](#commands)
- [Project Structure](#project-structure)
- [Healthchecks](#healthchecks)
- [Unit Test](#unittest)
- [About Graphql](#about-graphql)

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

1. Local mode

First, install golang dependency by using go module:

```bash
go mod tidy
```

then, run the app by using this command:

```bash
go run main.go
```

2. Docker Mode
```bash   
go mod tidy
docker build -t indraoctama/case-study-checkout:1.0.4 .
docker run -p 9000:9000 --rm indraoctama/case-study-checkout:1.0.4
```


webserver is going to running on `localhost:9000`.

## Project Structure
| Folder Name | Contain | Description |
|------|---------|-------------|
| constants |constanta| define static variable like : "PROMO_TYPE" |
| entity | data structure | organize data structure |
| handler | graphql or REST API handler | to handler HTTP request |
| pkg | helper and config | to manage local library and helper |
| repository | connect to datasource : checkout, products | to manage third-party repo such as database, partner, etc |
| usecase | business process | to manage business process |
| vendor | golang third-party library | will be showing after run command: `go mod tidy` |



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

## Unit Test

| File Name |  Description |
|------|-------------|
| checkout_case1_test.go| test case : promo macbook pro |
| checkout_case2_test.go | test case : promo google home |
| checkout_case3_test.go | test case : promo alexa, I have opinion about the question that the correct expected total is $394.2 |
| data_products_test.go | test case: get all and get one products |
| data_promo_test.go | test case: to test promo rule |

## About Graphql

I dont have experience about GraphQL but I want to try this:

1. To add cart
```cgo
curl --location --request POST 'http://localhost:9000/checkout/graphql' \
--header 'Content-Type: application/json' \
--data '{"query":"mutation { addCart(\n    sku:\"234234\",\n    qty:2\n    ) \n    { sku qty } \n}","variables":{}}'
```
