# Imagem base do Go
FROM golang:latest

# Instalar dependências para o Jenkins
RUN apt-get update && apt-get install -y wget gnupg

# Adicionar chave GPG do Jenkins e instalar o pacote do Jenkins
RUN wget -q -O - https://pkg.jenkins.io/debian/jenkins.io.key | apt-key add -
RUN sh -c 'echo deb http://pkg.jenkins.io/debian-stable binary/ > /etc/apt/sources.list.d/jenkins.list'
RUN apt-get update && apt-get install -y jenkins

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
