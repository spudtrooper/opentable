package main

import (
	"context"

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/opentable/cli"
)

func main() {
	check.Err(cli.Main(context.Background()))
}
