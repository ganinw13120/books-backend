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