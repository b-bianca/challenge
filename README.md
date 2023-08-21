# Meli Challenge

## Sumário
- [Sobre](#sobre)
- [Requisitos](#requisitos)
- [O que esse projeto faz e possui](#o-que-esse-projeto-faz-e-possui)
- [O que esse projeto não faz e débitos técnicos](#o-que-esse-projeto-não-faz-e-débitos-técnicos)
- [Como Executar o Projeto](#como-executar-o-projeto)
  - [Criar Variáveis de Ambiente](#criar-variáveis-de-ambiente)
  - [Executar o projeto](#como-executar-o-projeto)
  - [Executar o Docker](#executar-o-docker)
  - [Utilizar Aplicaçção & Documentação API](#utilizar-aplicação--documentação-api)
  - [Demonstração Rodando Docker Compose e Consumindo API](#demonstração-rodando-docker-compose-e-consumindo-api)


## Sobre
Este projeto visa simular um sistema de envio de notificações agendadas para usuários.

## Requisitos

|Recurso|Versão|Obrigatório|Nota|
|-|-|-|-|
|Docker Desktop| 4.21 ou mais atual|Sim|Necessário para rodar containers das APIs e banco de dados|
|Golang| 1.20|Não|Necessário apenas no caso de rodar localmente sem container|

## O que esse projeto faz e possui
### O que esse projeto faz
Através de APIs é possível criar um usuário, e com esse usuário agendar e  enviar uma mensagem. Essa mensagem é simulada, tendo seu registro persistido no banco de dados apenas para efeito de evidencia do funcionamento da rotina de agendamento e envio.

#### O que esse projeto possui
[x] Dockerfile e DockerCompose
[x] Documentação para Consumo das APIs
[x] Testes Automatizados
[x] Componentes
  [x] API para cadastro de mensagens e agendamento
  [x] API para notificação
  [x] Pooling
  [x] Banco de dados

## O que esse projeto não faz e débitos técnicos
#### O que esse projeto não faz
- Não envia propriamente as notificações, apenas simula o envio agendado e registra no banco de dados para evidencia do funcionamento da rotina;

#### Débitos técnicos
[] Remoção paramêtros *hard coded*, como portas das aplicações.
[] Estabelecer um intervalo mínimo entre o momento de criação da notificação e o agendamento dessa mesma notificação.
[] Teste do scheduller

## Como executar o projeto
### Criar Variáveis de Ambiente
Criar um arquivo nomedo como `.env` na raiz do projeto contendo os seguintes valores.
~~~bash
POSTGRES_USER="postuser"
POSTGRES_PASSWORD="postpass"
POSTGRES_DB="meli"
~~~
Notas: dada a natureza desse projeto, o arquivo ".env" já está na pasta raiz, assim como, intencionalmente, há valores ***hard coded*** no código.

### Executar o projeto
É possivel executar o projeto através do Mekefile, a partir da linha de comando:
~~~bash
make run-project
~~~
Notas: o comando deve ser efetuado na pasta raiz do projeto

### Executar o Docker
Para executar o projeto, é necessário ter o `Docker Desktop` instalado. Com isso será possível criar as instancias usando o comando `docker compose` via IDE ou linha de comando conforme a seguir:
~~~bash
docker compose -f "docker-compose.yml" up -d --build
~~~
Notas: o comando deve ser efetuado na pasta raiz do projeto

### Utilizar Aplicação & Documentação API
1. Crie um usuário `[POST] localhost:8080/api/v1/user` 
2. Com o ID do usuário, é possivel criar uma notificação com a mensagem e a data a ser enviada `[POST] http://localhost:8081/api/v1/:user_id`
3. Após a criação da notificação, é necessário aguardar a hora da mensagem agendada
4. É possivel verificar todas as notificações criadas `GET http://localhost:8081/api/v1/notification`
4. No momento agendado a mensagem é enviada, verifique no banco de dados para visualizar o envio `GET "http://localhost:8081/api/v1/notification/message"`
5. Caso o usuário não queira mais receber notificações, é possivel realizar o opt-out `PATCH http://localhost:8082/api/v1/user/:user_id `

A documentação está disponível via Postman com os casos de consumo.

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/13244098-734faa73-a2e1-4b8e-8faf-42abaec3f5c7?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D13244098-734faa73-a2e1-4b8e-8faf-42abaec3f5c7%26entityType%3Dcollection%26workspaceId%3D5e98eea6-1218-49b0-abb5-3b3c919df553)
