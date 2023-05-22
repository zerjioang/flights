package main

import (
	"fmt"
	"github.com/zerjioang/flights/cmd"
	"github.com/zerjioang/flights/embed"
)

// main is the CLI app main entry point
func main() {
	fmt.Println("Running version: ", embed.Version)
	cmd.Execute()
}
