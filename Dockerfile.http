FROM golang:latest AS builder
 
WORKDIR /app
 
COPY . .
 
RUN go mod download
 
RUN go build -o httpserver ./cmd/httpserver/main.go
 
EXPOSE 8080
 
ENTRYPOINT [ "/app/httpserver" ]