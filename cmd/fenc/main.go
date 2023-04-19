package main

import (
	"context"
	"flag"
	"os"

	"fenc/cmd/fenc/dec"
	"fenc/cmd/fenc/enc"
	"fenc/cmd/fenc/ver"

	"github.com/google/subcommands"
)

var (
	Version = "0.0.0"
)

// main
func main() {

	subcommands.Register(&enc.Encode{}, "")
	subcommands.Register(&dec.Decode{}, "")
	subcommands.Register(&ver.Ver{Semver: Version}, "")

	flag.Parse()

	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))

}
