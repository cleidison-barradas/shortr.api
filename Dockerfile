# ---- Stage 1: Build ----
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copia primeiro os arquivos de dependência (cache de layers)
COPY go.mod go.sum ./
RUN go mod download

# Copia o resto do código
COPY . .

# Build estático, sem CGO, otimizado
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s" \
    -o /app/server ./cmd/api

# ---- Stage 2: Runtime ----
FROM scratch

# Certificados TLS (necessário se seu app faz requests HTTPS)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080

USER 65534:65534

ENTRYPOINT ["/app/server"]