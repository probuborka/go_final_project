FROM golang:1.22 as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -o todo_lisl ./cmd/main.go

FROM ubuntu:latest as todo
ENV TODO_PORT=7540
ENV TODO_DBFILE=scheduler.db
ENV TODO_PASSWORD=test
WORKDIR /app
COPY --from=builder /app/todo_lisl /app/todo_lisl
COPY --from=builder /app/web /app/web
ENTRYPOINT ["/app/todo_lisl"]