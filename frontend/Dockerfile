# Usa uma imagem leve do Node
FROM node:20-alpine

# Define diretório de trabalho
WORKDIR /app

# Copia apenas os arquivos de dependência primeiro (para cache de build)
COPY package*.json ./

# Instala dependências
RUN npm install

# Copia o restante do código
COPY . .

# Expõe a porta padrão do Vite
EXPOSE 5173

# Comando padrão: iniciar servidor dev
CMD ["npm", "run", "dev", "--", "--host"]
