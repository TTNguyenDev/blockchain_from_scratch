// Package main is the entrypoint of this program
package main

import (
	"blockchain_from_scratch/cli"
)

func main() {
	cli := cli.CLI{}
	cli.Run()
}
