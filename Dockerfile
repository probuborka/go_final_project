#golang build todo-list
FROM golang:1.22.1-alpine as builder

ENV GOOS linux

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /app/todo_list ./cmd/todo/main.go

#alpine run todo-list
FROM alpine:latest as todo

ENV TODO_PORT=7540

ENV TODO_DBFILE=scheduler.db

ENV TODO_PASSWORD=test

WORKDIR /app

COPY --from=builder /app/todo_list /app/todo_list

COPY --from=builder /app/web /app/web

ENTRYPOINT ["/app/todo_list"]