# Imagem base do Go
FROM golang:latest

# Copiar os arquivos do código fonte
COPY producer.go /app/
COPY consumer.go /app/

# Definir o diretório de trabalho
WORKDIR /app

# Compilar os arquivos de código fonte
RUN go build producer.go
RUN go build consumer.go

# Expor as portas utilizadas pelo programa
EXPOSE 8080

# Comando para executar o programa
CMD ["./producer", "&", "./consumer"]
