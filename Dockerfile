# Imagen base
FROM golang:1.20 AS builder

# Directorio de trabajo
WORKDIR /app

# Copiar el código fuente
COPY . .

# Descargar dependencias y construir la aplicación
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp

# Crear una imagen más pequeña para la etapa de producción
FROM alpine:latest

# Copiar el binario del builder
COPY --from=builder /app/myapp /myapp

# Ejecutar la aplicación
CMD ["/myapp"]
