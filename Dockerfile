FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copiar archivos de módulos primero (para cache)
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente Y el .env
COPY . .
COPY .env .env

# Build de la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/.env .env
EXPOSE 8080
CMD ["./main"]