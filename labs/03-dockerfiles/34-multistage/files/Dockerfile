# ETAPA 1: Compilación (La llamamos "builder")
FROM golang:1.22-alpine AS builder

# Seteamos directorio de trabajo
WORKDIR /app

# Copiamos el código fuente
COPY main.go .

# Compilamos el binario
RUN go build -o mi-app-binario main.go


# ETAPA 2: Ejecución (Imagen final ultra ligera)
FROM alpine:latest

# Seteamos directorio de trabajo
WORKDIR /root/

# Copiamos SOLO el binario desde la etapa "builder"
COPY --from=builder /app/mi-app-binario .

# Exponemos el puerto y ejecutamos
EXPOSE 8080
CMD ["./mi-app-binario"]
