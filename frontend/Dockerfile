# Usar una imagen base de node
FROM node:16

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos del código fuente
COPY . .

# Instalar las dependencias
RUN npm install

# Construir el código
RUN npm run build

# Exponer el puerto 5173
EXPOSE 5173

# Servir los archivos estáticos con vite
CMD npm run dev