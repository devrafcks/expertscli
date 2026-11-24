package internal

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type Command interface {
	Name() string
	Example() string
	Help() string
	LongHelp() string
	Register(*flag.FlagSet)
	Run()
}

type CommandRoot struct {
	Name string
	commands []Command
}

func CommandInit(name string) *CommandRoot {
	return 	&CommandRoot{
		Name: name,
	}
}

func (cr *CommandRoot) Start(commandList []Command) error{
	if len(commandList) == 0{
		return errors.New("para inicializar o cli é solicitado um ou mais comandos")
	}
	cr.commands = commandList

	if len(os.Args) < 2 {
		//showHelp()
		return errors.New("erro: informe algum comando")
	}

	userPassedArguments := os.Args[1:]

	UserCommand := argumentFilter(userPassedArguments)

	fmt.Println(UserCommand)
	return nil
}
