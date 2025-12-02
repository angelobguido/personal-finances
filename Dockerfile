# Build stage para frontend
FROM node:20-alpine AS frontend-builder

WORKDIR /app/frontend

# Copiar package.json e instalar dependências
COPY frontend/package.json frontend/package-lock.json* ./
RUN npm install

# Copiar código fonte e buildar
COPY frontend/ ./
RUN npm run build

# Build stage para backend
FROM golang:1.24-alpine AS backend-builder

WORKDIR /app

# Copiar go.mod e go.sum
COPY go.mod go.sum* ./
RUN go mod download

# Copiar código fonte
COPY . .

# Buildar a aplicação Go
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server .

# Stage final
FROM alpine:latest

WORKDIR /app

# Instalar ca-certificates para conexões HTTPS/SSL
RUN apk --no-cache add ca-certificates

# Copiar o binário do backend
COPY --from=backend-builder /app/server /app/server

# Copiar arquivos estáticos
COPY --from=backend-builder /app/static /app/static

# Copiar o frontend buildado
COPY --from=frontend-builder /app/frontend/dist /app/frontend/dist

# Expor a porta
EXPOSE 5000

# Executar a aplicação
CMD ["/app/server"]
