package main

import (
	"fmt"

	"github.com/devrafcks/expertscli/cmd/commands"
	"github.com/devrafcks/expertscli/cmd/internal"
)

func main() {
	err := internal.CommandInit("expertscli").Start([]internal.Command{
		&commands.Start{},
		&commands.Message{},
		&commands.Ping{},
		&commands.Echo{},
		&commands.Version{},
		&commands.Env{},
		&commands.Hash{},
		&commands.Base64{},
		&commands.UUID{},
	})
	if err != nil {
		fmt.Println(err)
	}
}
