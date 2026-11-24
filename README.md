# expertscli

![Go Version](https://img.shields.io/badge/go-1.27.4-00ADD8?style=flat&logo=go&logoColor=white)
![Rocketseat](https://img.shields.io/badge/feito%20no-intensivo%20Rocketseat-8257e5?style=flat)

CLI construida em Go durante o intensivo de Golang da Rocketseat.

---

## Requisitos

- [Go](https://golang.org/dl/) 1.21+

---

## Instalacao

```bash
git clone https://github.com/devrafcks/expertscli.git
cd expertscli
go run main.go <comando> [flags]
```

---

## Comandos

### `start`

```bash
expertscli start --port 8081 --version 1.0.0
```

| Flag        | Tipo   | Obrigatorio |
|-------------|--------|-------------|
| `--port`    | int    | nao         |
| `--version` | string | sim         |

### `message`

```bash
expertscli message --text "Hello, world!"
```

| Flag     | Tipo   | Obrigatorio |
|----------|--------|-------------|
| `--text` | string | sim         |

> `*` not null

---

## Autor

Feito por [devrafcks](https://github.com/devrafcks) durante o intensivo de Go da Rocketseat.
