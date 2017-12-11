package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"

	"github.com/polydawn/rosetta"
	"github.com/polydawn/rosetta/cipher"
)

func main() {
	os.Exit(Main(os.Args, os.Stdin, os.Stdout, os.Stderr))
}

func Main(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	app := cli.NewApp()
	app.Name = "rosetta"
	app.Usage = "a tool for storing secrets with deterministic encryption."
	app.Commands = []cli.Command{
		cli.Command{
			Category: "basic",
			Name:     "encrypt",
			Usage:    "encrypt a stream fed to stdin, print to stdout",
			Action: func(c *cli.Context) error {
				key := cipher.Key{} // TODO cram key-gettery into helper method

				cleartext, err := ioutil.ReadAll(stdin)
				if err != nil && err != io.EOF {
					return err
				}

				enveloped, err := rosetta.EncryptAndEnvelopeBytes(cleartext, key)
				if err != nil {
					return err
				}

				_, err = stdout.Write(enveloped)
				return err
			},
		},
	}
	app.Writer = stdout
	app.ErrWriter = stderr
	err := app.Run(args)
	if err != nil {
		fmt.Fprintf(stderr, "error: %s\n", err)
		return 1
	}
	return 0
}
