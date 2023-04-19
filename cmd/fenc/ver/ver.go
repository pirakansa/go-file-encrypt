package ver

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

type Ver struct {
	Semver string
}

func (c *Ver) Name() string { return "ver" }

func (c *Ver) Synopsis() string { return "print version" }

func (c *Ver) Usage() string { return fmt.Sprintf("%s [args]", c.Name()) }

func (c *Ver) SetFlags(f *flag.FlagSet) {}

func (c *Ver) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	fmt.Printf("Version : %s\n", c.Semver)

	return subcommands.ExitSuccess
}
