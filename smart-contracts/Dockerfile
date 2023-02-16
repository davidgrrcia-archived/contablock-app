# Usar una imagen base de node
FROM node:16

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos del código fuente
COPY . .

# Instalar las dependencias
RUN npm install

# Expose the port used by Hardhat
EXPOSE 8545

# Set the command to run when the container starts
CMD npx hardhat node;

