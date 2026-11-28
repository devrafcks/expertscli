package commands

import (
	"crypto/rand"
	"flag"
	"fmt"
)

type UUID struct{}

const helpTextUUID = "gera um UUID v4 aleatorio"
const helpLongTextUUID = `Gera um UUID v4 criptograficamente aleatorio.

expertscli uuid`

const exampleTextUUID = `expertscli uuid`

func (cmd *UUID) Name() string {
	return "uuid"
}

func (cmd *UUID) Help() string {
	return helpTextUUID
}

func (cmd *UUID) LongHelp() string {
	return helpLongTextUUID
}

func (cmd *UUID) Example() string {
	return exampleTextUUID
}

func (cmd *UUID) Register(f *flag.FlagSet) {}

func (cmd *UUID) Run() {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Printf("erro ao gerar uuid: %s\n", err)
		return
	}

	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80

	fmt.Printf("%x-%x-%x-%x-%x\n", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
