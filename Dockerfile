FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum* ./

RUN go mod download

COPY . .

RUN go build -o /app/server /app/cmd/app

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/server .

RUN adduser -D -u 1000 apiuser
USER apiuser

EXPOSE 8080

ENV PORT=8080 \
    GIN_MODE=release

CMD ["./server"]