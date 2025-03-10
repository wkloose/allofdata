FROM golang:1.24.0 AS builder

WORKDIR /buiid

COPY . .

RUN CGO_ENABLED=0 go build -o trashure cmd/api/main.go

FROM alpine:latest AS prod

WORKDIR /app

COPY --from=builder /build/trashure .

EXPOSE 8080
ENTRYPOINT ["./trashure"]