# Usar una imagen base de golang
FROM golang:1.17

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos del código fuente
COPY . .

# Instalar las dependencias
RUN go mod download

# Compilar el código
RUN go build -o main .

# Exponer el puerto 8080
EXPOSE 8080

# Ejecutar el binario
CMD ["./main"]
