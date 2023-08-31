# Usa la imagen base de Go para compilar la aplicación
FROM golang:1.20.5-alpine3.17 as builder

# Establece el directorio de trabajo en el que se copiará el código
WORKDIR /app
RUN apk update && \
    apk add --no-cache gcc musl-dev

# Copia el archivo go.mod y go.sum al directorio de trabajo
COPY go.mod go.sum ./

# Descarga e instala las dependencias del módulo
RUN go mod download

# Copia el resto de los archivos al directorio de trabajo
COPY src/ ./src/

# Ejecutamos los test
RUN go test ./... -v

# Compila el código en un binario ejecutable
RUN CGO_ENABLED=1 GOOS=linux go build -o app ./src

# Crea una imagen mínima para ejecutar la aplicación
FROM alpine

# Copia el ejecutable de la aplicación desde la imagen del builder
COPY --from=builder /app/app /app/app

# Expone el puerto en el que se ejecuta la aplicación
EXPOSE 3000

# Ejecuta la aplicación cuando se inicie el contenedor
CMD ["/app/app"]