# 📌 Desafio Técnico Gerado pelo ChatGPT – API REST em Go

## 🎯 Objetivo

Construir uma API RESTful em Go que gerencie uma lista de tarefas (To-Do List).
O candidato deve implementar endpoints básicos, lidar com persistência e aplicar boas práticas de código.

<br>
<br>

## 🚀 Running

1. Run docker compose

```shell
docker compose up -d
```

2. Execute migrations

```shell
make migration-up
```

3. Run main.go

```shell
cd cmd/server && go run main.go
```

<br>
<br>

## 🚀 Build

```shell
docker compose up --build
```

<br>
<br>

## ⚙️ Requisitos Obrigatórios

1. CRUD de Tarefas
   - Criar uma tarefa
   - Listar todas as tarefas
   - Buscar uma tarefa por ID
   - Atualizar uma tarefa
   - Remover uma tarefa

Cada tarefa deve ter os seguintes campos:

```json
{
  "id": 1,
  "title": "Estudar Go",
  "description": "Ler documentação sobre goroutines",
  "done": false,
  "createdAt": "2025-09-24T10:00:00Z"
}
```

2. Validações

   - title é obrigatório.
   - description não pode ultrapassar 255 caracteres.

3. Persistência

   - Pode ser feito em memória (slice/map) ou utilizando um banco (SQLite/Postgres).
   - O candidato pode escolher a solução.

4. Boas práticas
   - Estruturar o projeto em camadas (ex: handlers, models, repository, routes).
   - Utilizar context onde for necessário.
   - Retornar respostas em JSON com status HTTP adequado.

<br>
<br>

## 🔥 Diferenciais (opcionais, mas contam pontos extras)

    - Implementar testes unitários ou de integração.
    - Documentar a API com Swagger/OpenAPI.
    - Utilizar Docker para subir o projeto.
    - Implementar paginação na listagem de tarefas.
    - Adicionar autenticação simples (ex: JWT ou token estático).

<br>
<br>

## 📡 Endpoints Esperados

| Método | Rota          | Descrição                      |
| ------ | ------------- | ------------------------------ |
| POST   | `/tasks`      | Criar uma nova tarefa          |
| GET    | `/tasks`      | Listar todas as tarefas        |
| GET    | `/tasks/{id}` | Buscar uma tarefa por ID       |
| PUT    | `/tasks/{id}` | Atualizar uma tarefa existente |
| DELETE | `/tasks/{id}` | Remover uma tarefa             |

<br>
<br>

## 📦 Entrega

    - Código hospedado em um repositório público no GitHub/GitLab.
    - Instruções claras no README.md de como rodar o projeto (com ou sem Docker).
    - Se possível, incluir exemplos de requisições no README.md (ex: via curl ou httpie).
