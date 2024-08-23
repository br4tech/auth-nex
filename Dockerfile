# Use a imagem oficial do Golang como base
FROM golang:1.21.0

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o arquivo go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixe as dependências do módulo Go
RUN go mod download

# Copie todo o código fonte para o diretório de trabalho
COPY . .

# Compile a aplicação Go para o sistema operacional Linux e arquitetura AMD64
RUN go build -o app cmd/main.go

# Exponha a porta 8080 para o mundo externo
EXPOSE 8080

# Comando para executar a aplicação quando o contêiner for iniciado
CMD ["./app"]