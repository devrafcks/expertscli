package commands

import (
	"flag"
	"fmt"
)

type Version struct{}

const helpTextVersion = "exibe a versao do expertscli"
const helpLongTextVersion = `Exibe informacoes sobre a versao atual do expertscli.

expertscli version`

const exampleTextVersion = `expertscli version`

func (cmd *Version) Name() string {
	return "version"
}

func (cmd *Version) Help() string {
	return helpTextVersion
}

func (cmd *Version) LongHelp() string {
	return helpLongTextVersion
}

func (cmd *Version) Example() string {
	return exampleTextVersion
}

func (cmd *Version) Register(f *flag.FlagSet) {}

func (cmd *Version) Run() {
	fmt.Println("expertscli v1.0.0")
	fmt.Println("go intensivo - rocketseat")
}
