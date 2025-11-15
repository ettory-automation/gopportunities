# =====================
# STAGE 1 — BUILD
# =====================
FROM golang:1.25.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Install dependencies to CGO + sqlite3
RUN apt-get update && apt-get install -y gcc libc6-dev libsqlite3-dev

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o gopportunities main.go

# =====================
# STAGE 2 — RUNTIME
# =====================
FROM debian:bookworm-slim

WORKDIR /app

RUN apt-get update && apt-get install -y libsqlite3-0

COPY --from=builder /app/gopportunities .

EXPOSE 8080

CMD ["./gopportunities"]
