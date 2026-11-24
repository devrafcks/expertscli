package main

import (
	"fmt"
	"github.com/devrafcks/expertscli/cmd/internal"
)

func main() {
	err := internal.CommandInit("expertscli").Start([]internal.Command{})
	if err != nil {
		fmt.Println(err)
	}
}