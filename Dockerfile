FROM golang:latest AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o monitoring cmd/app.go

FROM ubuntu:latest
RUN apt update; apt upgrade -y
RUN apt install -y ca-certificates libc6
WORKDIR /app
COPY --from=build /app/monitoring monitoring
ENTRYPOINT ["/app/monitoring"]
