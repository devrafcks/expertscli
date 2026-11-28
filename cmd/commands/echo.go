package commands

import (
	"flag"
	"fmt"
	"strings"
)

type Echo struct {
	Text  string
	Upper bool
}

const helpTextEcho = "ecoa um texto no terminal"
const helpLongTextEcho = `Exibe o texto informado no terminal. Use --upper para converter em maiusculas.

expertscli echo --text [text] [--upper]
  - text: string
  - upper: bool`

const exampleTextEcho = `expertscli echo --text "Hello, world!"
expertscli echo --text "hello" --upper`

func (cmd *Echo) Name() string {
	return "echo"
}

func (cmd *Echo) Help() string {
	return helpTextEcho
}

func (cmd *Echo) LongHelp() string {
	return helpLongTextEcho
}

func (cmd *Echo) Example() string {
	return exampleTextEcho
}

func (cmd *Echo) Register(f *flag.FlagSet) {
	f.StringVar(&cmd.Text, "text", "", "texto para ecoar")
	f.BoolVar(&cmd.Upper, "upper", false, "converte para maiusculo")
}

func (cmd *Echo) Run() {
	if cmd.Text == "" {
		fmt.Println("erro: informe um texto com --text")
		return
	}

	text := cmd.Text
	if cmd.Upper {
		text = strings.ToUpper(text)
	}

	fmt.Println(text)
}
