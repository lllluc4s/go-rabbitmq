# Imagem base do Go
FROM golang:latest

# Instalar dependências
RUN apt-get update && \
    apt-get install -y git vim && \
    rm -rf /var/lib/apt/lists/*

# Definir variável de ambiente para o PATH
ENV PATH="$PATH:/usr/local/go/bin"

# Copiar os arquivos do código fonte
COPY . /app

# Definir o diretório de trabalho
WORKDIR /app

# Rodar o comando go mod tidy para garantir que as dependências estão corretas
RUN go mod tidy

# Compilar os arquivos de código fonte
RUN go build producer.go
RUN go build consumer.go

# Expor as portas utilizadas pelo programa
EXPOSE 8080

# Comando para executar o programa
CMD ["sh", "-c", "./producer & ./consumer"]
