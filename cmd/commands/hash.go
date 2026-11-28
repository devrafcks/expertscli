package commands

import (
	"crypto/md5"
	"crypto/sha256"
	"flag"
	"fmt"
)

type Hash struct {
	Text string
	Algo string
}

const helpTextHash = "gera o hash de um texto"
const helpLongTextHash = `Gera o hash SHA256 ou MD5 de um texto informado.

expertscli hash --text [text] [--algo sha256|md5]
  - text: string
  - algo: string (default: sha256)`

const exampleTextHash = `expertscli hash --text "Hello, world!"
expertscli hash --text "Hello, world!" --algo md5`

func (cmd *Hash) Name() string {
	return "hash"
}

func (cmd *Hash) Help() string {
	return helpTextHash
}

func (cmd *Hash) LongHelp() string {
	return helpLongTextHash
}

func (cmd *Hash) Example() string {
	return exampleTextHash
}

func (cmd *Hash) Register(f *flag.FlagSet) {
	f.StringVar(&cmd.Text, "text", "", "texto para gerar o hash")
	f.StringVar(&cmd.Algo, "algo", "sha256", "algoritmo: sha256 ou md5")
}

func (cmd *Hash) Run() {
	if cmd.Text == "" {
		fmt.Println("erro: informe um texto com --text")
		return
	}

	switch cmd.Algo {
	case "md5":
		h := md5.Sum([]byte(cmd.Text))
		fmt.Printf("md5: %x\n", h)
	case "sha256":
		h := sha256.Sum256([]byte(cmd.Text))
		fmt.Printf("sha256: %x\n", h)
	default:
		fmt.Printf("erro: algoritmo invalido '%s'. Use sha256 ou md5\n", cmd.Algo)
	}
}
