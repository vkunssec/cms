FROM golang:1.20 as build

WORKDIR /app

COPY . /app/

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o application ./cmd/api

FROM alpine:3.9 as deploySession

RUN apk add ca-certificates

WORKDIR /app

COPY --from=build /app /app

CMD ["./application"]
