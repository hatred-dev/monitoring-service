FROM golang:bookworm AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -o monitoring cmd/app.go

FROM debian:bookworm
RUN apt update; apt install -y ca-certificates
WORKDIR /app
COPY --from=build /app/monitoring monitoring
ENTRYPOINT ["/app/monitoring"]
