package dec

import (
	"compress/gzip"
	"context"
	"crypto/cipher"
	"fenc/internal/encdata"
	"fenc/internal/input"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/google/subcommands"
)

type Decode struct {
	ifpath string
	ofpath string
	pass   string
}

func (c *Decode) Name() string { return "dec" }

func (c *Decode) Synopsis() string { return "do decode" }

func (c *Decode) Usage() string { return fmt.Sprintf("Usage: %s -i <path> -o <path>\n", c.Name()) }

func (c *Decode) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.ifpath, "i", "", "input file")
	f.StringVar(&c.ofpath, "o", "", "output file")
	f.StringVar(&c.pass, "p", "", "noninteract pass")
}

func (c *Decode) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	infile, err := os.OpenFile(c.ifpath, os.O_RDONLY, 0755)
	if err != nil {
		fmt.Printf("Err: %s\n", err.Error())
		f.Usage()
		os.Exit(int(subcommands.ExitFailure))
	}
	defer infile.Close()

	outfile, err := os.OpenFile(c.ofpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_EXCL, 0755)
	if err != nil {
		fmt.Printf("Err: %s\n", err.Error())
		f.Usage()
		os.Exit(int(subcommands.ExitFailure))
	}
	defer outfile.Close()

	greader, err := gzip.NewReader(infile)
	if err != nil {
		panic(err)
	}
	defer greader.Close()

	pass := c.pass
	if pass == "" {
		pass, err = input.ReadPassword()
		if err != nil {
			fmt.Printf("Err: %s", err.Error())
			os.Exit(int(subcommands.ExitFailure))
		}
	}

	key := encdata.GetEncriptKey(pass)
	stream := encdata.GetDecipherObj(key)

	writer := cipher.StreamWriter{S: stream, W: outfile}
	io.Copy(writer, greader)

	return subcommands.ExitSuccess
}
