package commands

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

type Ping struct {
	URL string
}

const helpTextPing = "verifica se uma url esta acessivel"
const helpLongTextPing = `Faz uma requisicao HTTP GET para a URL informada e retorna o status e o tempo de resposta.

expertscli ping --url [url]
  - url: string`

const exampleTextPing = `expertscli ping --url https://google.com
expertscli ping --url https://api.exemplo.com`

func (cmd *Ping) Name() string {
	return "ping"
}

func (cmd *Ping) Help() string {
	return helpTextPing
}

func (cmd *Ping) LongHelp() string {
	return helpLongTextPing
}

func (cmd *Ping) Example() string {
	return exampleTextPing
}

func (cmd *Ping) Register(f *flag.FlagSet) {
	f.StringVar(&cmd.URL, "url", "", "url para verificar")
}

func (cmd *Ping) Run() {
	if cmd.URL == "" {
		fmt.Println("erro: informe uma url com --url")
		return
	}

	fmt.Printf("verificando %s...\n", cmd.URL)

	start := time.Now()
	resp, err := http.Get(cmd.URL)
	if err != nil {
		fmt.Printf("erro ao conectar: %s\n", err)
		return
	}
	defer resp.Body.Close()

	elapsed := time.Since(start).Round(time.Millisecond)
	fmt.Printf("status: %d | tempo: %s\n", resp.StatusCode, elapsed)
}
