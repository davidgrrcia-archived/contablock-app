# Usar una imagen base de golang
FROM golang:1.17

# Copy backend library
RUN mkdir ../backend
COPY . .

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos del código fuente
COPY . .

# Instalar las dependencias
RUN go mod download

# Compilar el código
RUN go build -o main .

# Exponer el puerto 3000
EXPOSE 3000

# Ejecutar el binario
CMD ["./main"]
