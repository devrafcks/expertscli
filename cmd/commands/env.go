package commands

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Env struct {
	Filter string
}

const helpTextEnv = "lista as variaveis de ambiente do sistema"
const helpLongTextEnv = `Lista todas as variaveis de ambiente. Use --filter para buscar por nome.

expertscli env [--filter nome]
  - filter: string`

const exampleTextEnv = `expertscli env
expertscli env --filter PATH`

func (cmd *Env) Name() string {
	return "env"
}

func (cmd *Env) Help() string {
	return helpTextEnv
}

func (cmd *Env) LongHelp() string {
	return helpLongTextEnv
}

func (cmd *Env) Example() string {
	return exampleTextEnv
}

func (cmd *Env) Register(f *flag.FlagSet) {
	f.StringVar(&cmd.Filter, "filter", "", "filtra variaveis pelo nome")
}

func (cmd *Env) Run() {
	envs := os.Environ()
	filter := strings.ToUpper(cmd.Filter)
	count := 0

	for _, e := range envs {
		if filter == "" || strings.Contains(strings.ToUpper(e), filter) {
			fmt.Println(e)
			count++
		}
	}

	fmt.Printf("\n%d variavel(is) encontrada(s)\n", count)
}
