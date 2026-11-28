package commands

import (
	"flag"
	"fmt"
)

type Start struct {
	Port    int
	Version string
}

const helpTextStart = "responsavel por inicializar o servidor"
const helpLongTextStart = `Responsável por inicializar um servidor simples liberando acesso a todos os serviços configurados.

expertscli start --port [serverPort] --version [serverVersion]
  - serverPort: int
  - serverVersion: string`

const exampleTextStart = `expertscli start --port 3030 --version 1.0.0
expertscli start --version 1.0.0 ...`

func (cmd *Start) Name() string {
	return "start"
}

func (cmd *Start) Help() string {
	return helpTextStart
}

func (cmd *Start) LongHelp() string {
	return helpLongTextStart
}

func (cmd *Start) Example() string {
	return exampleTextStart
}

func (cmd *Start) Register(f *flag.FlagSet) {
	f.IntVar(&cmd.Port, "port", 3000, "porta do servidor")
	f.StringVar(&cmd.Version, "version", "", "versao do servidor")
}

func (cmd *Start) Run() {
	fmt.Printf("servidor iniciado | porta: %d | versao: %s\n", cmd.Port, cmd.Version)
}
