# Etapa de compilação
FROM golang:1.21 AS builder

WORKDIR /app

# Copiar os arquivos de gerenciamento de dependências e baixá-los
COPY go.mod go.sum ./
RUN go mod download

# Copiar o restante dos arquivos fonte e construir o aplicativo
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /ms-ddconnector .

# Etapa final da imagem
FROM alpine:3.19.1
RUN apk --no-cache add ca-certificates && \
    apk --no-cache upgrade && \
    apk add openssl=3.1.4-r6

# Definir o diretório de trabalho
WORKDIR /root/

# Copiar o binário da etapa de compilação para a imagem final
COPY --from=builder /ms-ddconnector .

# Definir o comando de execução
ENTRYPOINT ["./ms-ddconnector"]

