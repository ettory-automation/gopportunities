<div style="display: flex; align-items: center; justify-content: center; background-color: #4b3c8a; padding: 20px; border-radius: 10px;">
  <img alt="Go" height="60" width="80" src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" style="margin: 0 15px;">
  <img alt="Gin" height="60" width="80" src="https://gin-gonic.com/_astro/gin.D6H2T_2v_ZD2G7l.webp" style="margin: 0 15px;">
  <img alt="SQLite" height="60" width="80" src="https://raw.githubusercontent.com/devicons/devicon/master/icons/sqlite/sqlite-original.svg" style="margin: 0 15px;">
  <img alt="Docker" height="60" width="80" src="https://raw.githubusercontent.com/devicons/devicon/master/icons/docker/docker-original.svg" style="margin: 0 15px;">
  <img alt="Caddy" height="60" width="80" src="https://caddyserver.com/old/resources/images/caddy-logo.svg" style="margin: 0 15px;">
  <img alt="Swagger" height="60" width="80" src="https://raw.githubusercontent.com/devicons/devicon/master/icons/swagger/swagger-original.svg" style="margin: 0 15px;">
</div>

---

# GOpportunities

- Descrição: GOpportunities é uma API RESTful construída em Go para gerenciar oportunidades de emprego. O projeto utiliza Gin para rotas, SQLite como banco de dados e Swagger para documentação. O deploy é feito via Docker, com Caddy como reverse proxy.

---

1. Tecnologias utilizadas:
  - [Go 1.25](https://golang.org/)
  - [Gin](https://github.com/gin-gonic/gin)
  - [SQLite3](https://www.sqlite.org/index.html)
  - [Swaggo](https://github.com/swaggo/gin-swagger)
  - [Docker & Docker Compose](https://www.docker.com/)
  - [Caddy](https://caddyserver.com/)

2. Endpoints:

  - Healthcheck:
    GET /healthz
    Resposta: {"status": "ok"}

3. API de oportunidades de emprego (base /api/v1):

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| POST   | /opening | Cadastra uma nova oportunidade |
| GET    | /opening | Mostra uma oportunidade específica |
| PUT    | /opening | Atualiza uma oportunidade existente |
| DELETE | /opening | Remove uma oportunidade cadastrada|
| GET    | /openings | Lista todas as oportunidades cadastradas |

### Swagger
- Documentação automática da API disponível em: GET /swagger/index.html

## Rodando localmente:

1. Clone o repositório:

```bash
   git clone https://github.com/ettory-automation/gopportunities.git
   cd gopportunities
```

2. Build da imagem Docker:

```bash
   docker compose build
```

3. Suba os containers:

```bash
   docker compose up -d
```

## Acesso à documentação

- Acesse a documentação da API em Swagger no link: http://localhost:9052/swagger/index.html

## Estrutura do Projeto

- Docker:
  - Serviço 'gopportunities' na porta interna 8080, exposto via Caddy na porta externa 9052.
  - Dados persistidos no volume 'gopportunities_data'.
  - Healthcheck configurado no Dockerfile.

```plaintext
.
├── Caddyfile
├── config
│   ├── config.go
│   ├── logger.go
│   └── sqlite.go
├── docker-compose.yml
├── Dockerfile
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── handler
│   ├── createOpening.go
│   ├── deleteOpening.go
│   ├── handler.go
│   ├── listOpenings.go
│   ├── request.go
│   ├── response.go
│   ├── showOpening.go
│   └── updateOpening.go
├── main.go
├── makefile
├── router
│   ├── router.go
│   └── routes.go
└── schemas
    └── opening.go
```
