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
WORKDIR /app
COPY migrations migrations
COPY --from=build /app/monitoring monitoring
ENTRYPOINT ["/app/monitoring"]
