package main

import (
	"context"

	"github.com/spudtrooper/goutil/check"
	cli "github.com/spudtrooper/opentable/clinew"
)

func main() {
	check.Err(cli.Main(context.Background()))
}
