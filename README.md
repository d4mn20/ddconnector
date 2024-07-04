# Microserviço "ms-ddconnector"
O microserviço "ms-ddconnector" é responsável por receber e armazenar resultados de varreduras de segurança (SAST, SCA e DAST), que são enviados por meio de requisições advindas do Azure Pipelines. Ele faz parte de um sistema maior que integra varreduras de segurança em um ambiente DevSecOps e se conecta ao DefectDojo e ao Azure Copilot para processar e gerenciar os resultados.

# Funcionalidades
Recebe dados de varreduras de segurança via requisições HTTP do Azure Pipelines.
Armazena os resultados das varreduras em um banco de dados para processamento no DefectDojo.


# Configuração
Defina a porta em que o microserviço será executado no arquivo <router.go> e certifique-se de que ele seja acessível pela rede.

# Uso
Para usar o microserviço "ms-ddconnector", siga estas etapas:

## Clone este repositório:
```bash
git clone https://bbts-lab@dev.azure.com/bbts-lab/DevSecOps/_git/ms-ddconnector
```

Configure as variáveis de ambiente e os arquivos de configuração de acordo com a sua implementação.

Execute o microserviço:
```go
go run main.go
```

Envie requisições HTTP contendo os resultados das varreduras para o endpoint apropriado.

O microserviço processará as requisições e armazenará os resultados no banco de dados configurado.

## Para uso com Docker
Para construir e executar o container Docker do "ms-ddconnector", utilize os seguintes comandos:

```bash
sudo docker build -t ms-ddconnectori .
sudo docker run --name ms-ddconnector -d -e DD_API_KEY=<API_KEY> -p 21777:21777 ms-ddconnectori
```

# Licença
Este microserviço foi produzido pela BBTS e está licenciado sob a Licença MIT.

# To Do
- [ ] Adicionar "Field" de "Environment" no import e reimport.