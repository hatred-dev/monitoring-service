FROM golang:1.20-alpine as build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY src src
COPY main.go ./
RUN go build -o monitoring

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/monitoring monitoring
ENTRYPOINT ["/app/monitoring"]
