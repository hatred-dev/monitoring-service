FROM golang:latest as build
WORKDIR /app
COPY database database
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY src src
COPY main.go ./
RUN go build -o monitoring

FROM debian:bullseye-slim
RUN apt update; apt upgrade -y
RUN apt install -y ca-certificates
WORKDIR /app
COPY migrations migrations
COPY --from=build /app/monitoring monitoring
ENTRYPOINT ["/app/monitoring"]
