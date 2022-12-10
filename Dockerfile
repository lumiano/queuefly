ARG VERSION=1.19.3
ARG PORT=3000

FROM golang:${VERSION} AS builder

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/queuefly ./cmd/main.go

FROM scratch as production

COPY --from=builder /app/bin/queuefly ./queuefly
COPY .env .

EXPOSE $PORT

CMD ["./queuefly", "app:serve"]
