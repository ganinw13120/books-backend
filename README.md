# Books review backend service
> Get data from database, being a web service.

> Practice projects implementing with golang. Applying with hexagonal architecture.

Check out my frontend repository : https://github.com/ganinw13120/books-web

## Installation

```sh
go get
```

## Development

```sh
go run server.go
```

## Deployment

```sh
go build .
```

Deployment is done by Google Cloud Platform, including GCP Cloud Build, GCP Cloud Run, and GCP Cloud Storage (storing .env file).

Dockerfile :

```
FROM golang:1.16 AS builder

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 go build -o books-backend .

FROM alpine:3.13

RUN apk add tzdata
RUN cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime
RUN echo "Asia/Bangkok" >  /etc/timezone

WORKDIR /usr/src/app

COPY --from=builder /src/books-backend /usr/src/app/books-backend
COPY --from=builder /src/.env /usr/src/app/.env

EXPOSE 8080
CMD ["./books-backend"]
```
