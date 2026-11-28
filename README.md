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
go run ./cmd/expertscli/main.go <comando> [flags]
```

---

## Comandos

### `start`

Inicializa um servidor simples liberando acesso a todos os servicos configurados.

```bash
go run ./cmd/expertscli/main.go start --port 8080 --version 1.0.0
```

| Flag        | Tipo   | Obrigatorio | Default |
|-------------|--------|-------------|---------|
| `--port`    | int    | nao         | 3000    |
| `--version` | string | sim         | -       |

---

### `message`

Envia uma mensagem via CLI.

```bash
go run ./cmd/expertscli/main.go message --text "Hello, world!"
```

| Flag     | Tipo   | Obrigatorio |
|----------|--------|-------------|
| `--text` | string | sim         |

---

### `ping`

Verifica se uma URL esta acessivel, retornando status HTTP e tempo de resposta.

```bash
go run ./cmd/expertscli/main.go ping --url https://google.com
```

| Flag    | Tipo   | Obrigatorio |
|---------|--------|-------------|
| `--url` | string | sim         |

---

### `echo`

Ecoa um texto no terminal. Use `--upper` para converter em maiusculas.

```bash
go run ./cmd/expertscli/main.go echo --text "hello" --upper
```

| Flag      | Tipo   | Obrigatorio | Default |
|-----------|--------|-------------|---------|
| `--text`  | string | sim         | -       |
| `--upper` | bool   | nao         | false   |

---

### `version`

Exibe a versao atual do expertscli.

```bash
go run ./cmd/expertscli/main.go version
```

---

### `help`

Exibe a lista de todos os comandos disponiveis com exemplos de uso.

```bash
go run ./cmd/expertscli/main.go help
```

> `*` not null

---

## Estrutura

```
expertscli/
├── cmd/
│   ├── commands/
│   │   ├── start.go      # inicia um servidor
│   │   ├── message.go    # envia uma mensagem
│   │   ├── ping.go       # verifica uma url
│   │   ├── echo.go       # ecoa texto no terminal
│   │   └── version.go    # exibe a versao do cli
│   ├── expertscli/
│   │   └── main.go       # entrypoint
│   └── internal/
│       ├── command.go    # engine da cli
│       └── arguments.go  # parser de argumentos
└── go.mod
```

