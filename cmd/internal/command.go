package internal

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
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
		cr.showHelp()
		return errors.New("erro: informe algum comando")
	}

	userPassedArguments := os.Args[1:]

	UserCommand := argumentFilter(userPassedArguments)

		if UserCommand.Command == "" {
			cr.showHelp()
			return errors.New("Informe um comando valido")
		}

		if UserCommand.Command == "help" {
			cr.showHelp()
			return nil
		}


		for _, command := range cr.commands {
			if UserCommand.Command == command.Name(){
					flagset := flag.NewFlagSet(command.Name(), flag.ContinueOnError)
					command.Register(flagset)
					flagset.Parse(os.Args[2:])
					command.Run()
					return nil

			}
		}

		cr.showHelp()
	
	return fmt.Errorf("%s comando invalido", UserCommand.Command)
}

func (cr *CommandRoot) showHelp() {
	fmt.Printf("Usage: %s [COMMAND] [OPTIONS]\n\n", cr.Name)

	tabWriter := tabwriter.NewWriter(os.Stdout, 0,0,2, ' ' ,0) 

	fmt.Fprintf(tabWriter, "commands:\n\n")
	for _, command := range cr.commands {			
		fmt.Fprintf(tabWriter, "\t- %s\t%s\n", command.Name(), command.Help())	
	}
	tabWriter.Flush()
	fmt.Fprintf(tabWriter, "\nexamples:\n")
	for _, command := range cr.commands {			
		fmt.Fprintf(tabWriter, "\t%s\n", command.Example())	
	}

}