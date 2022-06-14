
# Simple Go Chat

![fluxo de trabalho de exemplo](https://github.com/wblech/chatGo/actions/workflows/actions.yml/badge.svg)


This is a very simple go chat using WebSockets and rabbitmq for getting stock data.
## Install

Install the project using Makefile
It will run a docker-compose up command. For installing keycloak, rabbitmq and database

```bash
  make start_dump
```

You need to wait Keycloak create the client. This might take 1 minute or 2 minutes

After that you can run

```bash
  go run src/main.go
```

If you get this error:

`Get "http://localhost:8081/auth/realms/chat/.well-known/openid-configuration": dial tcp 127.0.0.1:8081: connect: connection refused
`

It is because Keycloak is still creating the client. Go grab a coffee :)

    