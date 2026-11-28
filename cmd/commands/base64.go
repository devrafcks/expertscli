package commands

import (
	"encoding/base64"
	"flag"
	"fmt"
)

type Base64 struct {
	Text   string
	Decode bool
}

const helpTextBase64 = "encoda ou decoda um texto em base64"
const helpLongTextBase64 = `Encoda um texto para base64 ou decoda de base64. Use --decode para decodificar.

expertscli base64 --text [text] [--decode]
  - text: string
  - decode: bool`

const exampleTextBase64 = `expertscli base64 --text "Hello, world!"
expertscli base64 --text "SGVsbG8sIHdvcmxkIQ==" --decode`

func (cmd *Base64) Name() string {
	return "base64"
}

func (cmd *Base64) Help() string {
	return helpTextBase64
}

func (cmd *Base64) LongHelp() string {
	return helpLongTextBase64
}

func (cmd *Base64) Example() string {
	return exampleTextBase64
}

func (cmd *Base64) Register(f *flag.FlagSet) {
	f.StringVar(&cmd.Text, "text", "", "texto para encodar ou decodar")
	f.BoolVar(&cmd.Decode, "decode", false, "decoda ao inves de encodar")
}

func (cmd *Base64) Run() {
	if cmd.Text == "" {
		fmt.Println("erro: informe um texto com --text")
		return
	}

	if cmd.Decode {
		decoded, err := base64.StdEncoding.DecodeString(cmd.Text)
		if err != nil {
			fmt.Printf("erro ao decodar: %s\n", err)
			return
		}
		fmt.Println(string(decoded))
		return
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(cmd.Text))
	fmt.Println(encoded)
}
