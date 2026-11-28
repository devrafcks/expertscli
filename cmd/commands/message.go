package commands

import (
	"flag"
	"fmt"
)

type Message struct {
	Text string
}

const helpTextMessage = "envia uma mensagem via cli"
const helpLongTextMessage = `Responsável por enviar uma mensagem simples pelo terminal.

expertscli message --text [text]
  - text: string`

const exampleTextMessage = `expertscli message --text "Hello, world!"
expertscli message --text "Ola expertscli"`

func (cmd *Message) Name() string {
	return "message"
}

func (cmd *Message) Help() string {
	return helpTextMessage
}

func (cmd *Message) LongHelp() string {
	return helpLongTextMessage
}

func (cmd *Message) Example() string {
	return exampleTextMessage
}

func (cmd *Message) Register(f *flag.FlagSet) {
	f.StringVar(&cmd.Text, "text", "", "texto da mensagem")
}

func (cmd *Message) Run() {
	fmt.Printf("mensagem: %s\n", cmd.Text)
}
