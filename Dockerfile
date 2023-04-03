FROM golang:alpine as build
WORKDIR /app
COPY database database
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY src src
COPY main.go ./
RUN go build -o monitoring

FROM alpine:latest
WORKDIR /app
COPY migrations migrations
COPY --from=build /app/monitoring monitoring
ENTRYPOINT ["/app/monitoring"]
