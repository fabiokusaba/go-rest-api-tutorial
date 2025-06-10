# API REST em Go - Autenticação de Usuários

Este é um projeto de estudos focado na criação de uma API REST utilizando a linguagem Go (Golang) para implementar um sistema de registro e autenticação de usuários.

## 🎯 Objetivo

O objetivo principal deste projeto é aprender e demonstrar as melhores práticas no desenvolvimento de APIs REST em Go, incluindo:

- Estruturação de projetos Go
- Implementação de endpoints REST
- Autenticação e autorização de usuários
- Gerenciamento de banco de dados com migrations
- Segurança em APIs

## 🛠️ Tecnologias Utilizadas

- **Go (Golang)** - Linguagem de programação principal
- **PostgreSQL** - Banco de dados relacional
- **Migrations** - Controle de versão do banco de dados
- **Bcrypt** - Hash de senhas

## 📋 Funcionalidades

### Autenticação
- [x] Registro de novos usuários
- [x] Login de usuários existentes
- [x] Hash seguro de senhas

### Gerenciamento de Usuários
- [x] Criação de usuários
- [x] Validação de dados de entrada

## 🚀 Como Executar

### Pré-requisitos
- Go 1.24
- PostgreSQL 16
- Make (opcional)

### Configuração do Ambiente

1. Clone o repositório:
```bash
git clone <https://github.com/fabiokusaba/go-rest-api-tutorial.git>
cd go-rest-api-tutorial
```

2. Configure as variáveis de ambiente:
```bash
DB_HOST="localhost"
DB_PORT="5432"
DB_USER="postgres"
DB_PASSWORD="your-password"
DB_NAME="your-database-name"
# Edite o arquivo .env com suas configurações
```

3. Execute as migrations:
```bash
make migrate-up
# ou
go run cmd/migrate/main.go up
```

5. Instale as dependências:
```bash
go mod download
```

6. Execute a aplicação:
```bash
make run
# ou
go run cmd/main.go
```

A API estará disponível em `http://localhost:8080`

## 📚 Endpoints da API

```
POST /register - Registrar novo usuário
POST /login    - Login de usuário
```

## 📖 Migrations

### Criar nova migration:
```bash
make migration <name>
```

### Executar migrations:
```bash
make migrate-up
```

### Reverter migrations:
```bash
make migrate-down
```

## 🔐 Segurança

Este projeto implementa práticas de segurança:

- Senhas hasheadas com bcrypt
- Validação de entrada de dados

## 📚 Recursos de Aprendizado

- [Documentação oficial do Go](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
