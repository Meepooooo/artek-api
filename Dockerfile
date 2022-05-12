# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY main.go ./
COPY api/*.go ./api/
COPY config/config.go ./config/
COPY db/db.go ./db/

RUN go build -o /artek-api

ENV DB_LOCATION /tmp/app.db
ENV PORT 8080

EXPOSE 8080

CMD [ "/artek-api" ]