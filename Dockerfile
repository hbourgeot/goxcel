# Etapa 1: Construir el backend de Go
ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as go-builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app ./cmd/goxcel

# Etapa 2: Construir el frontend con Node.js y Vite
FROM node:18 as node-builder

WORKDIR /usr/src/app/frontend
COPY frontend/package.json frontend/yarn.lock ./
RUN yarn install

COPY frontend/ .
RUN yarn build

# Etapa 3: Imagen final que contiene tanto el frontend como el backend
FROM debian:bookworm

# Copiar el ejecutable de Go desde la etapa de construcción de Go
COPY --from=go-builder /run-app /usr/local/bin/

# Copiar los archivos estáticos construidos desde la etapa de construcción de Node
COPY --from=node-builder /usr/src/app/frontend/dist /usr/local/bin/frontend/dist

# Configura el directorio para servir los archivos estáticos del frontend
ENV FRONTEND_DIR=/usr/local/bin/frontend/dist

CMD ["run-app"]
