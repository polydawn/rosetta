package main

import (
	"encoding/base64"
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
	app.Description = "Rosetta is a simple, scriptable file encryption tool." +
		"\n" +
		"\n   Rosetta works with symmetric keys, saves your encrypted secrets" +
		"\n   in the standard PEM (RFC 1421) plaintext format for ease of use" +
		"\n   and simple copy-paste interaction, and includes helpful features" +
		"\n   like built-in password-derived key generation." +
		"\n" +
		"\n   Rosetta is designed to produce encryption that is *deterministic*" +
		"\n   from the content and keys.  This means Rosetta is can be used" +
		"\n   efficiently for storing secrets in e.g. a git repository; it" +
		"\n   will only generate diffs when the content is changed.  Go ahead." +
		"\n   Commit your ciphertext."
	app.Commands = []cli.Command{
		cli.Command{
			Category: "basic",
			Name:     "encrypt",
			Usage:    "encrypt a stream fed to stdin, print to stdout",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "p",
					Usage: "password (will be strengthed into key)",
				},
				cli.StringFlag{
					Name:  "k",
					Usage: "key (in base64)",
				},
			},
			Action: func(c *cli.Context) error {
				key, err := processKeyArgs(c)
				if err != nil {
					return err
				}

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

func processKeyArgs(c *cli.Context) (cipher.Key, error) {
	switch {
	case c.IsSet("p"):
		if c.IsSet("k") {
			return nil, fmt.Errorf("EITHER the '-k' or '-p' flag must be provided, not both")
		}
		password := c.String("p")
		// Key length param should vary based on suite, but since we only have the one right now...
		return rosetta.DeriveKey([]byte(password), 32)
	case c.IsSet("k"):
		if c.IsSet("p") {
			return nil, fmt.Errorf("EITHER the '-k' or '-p' flag must be provided, not both")
		}
		return base64.StdEncoding.DecodeString(c.String("k"))
	default:
		return nil, fmt.Errorf("either the '-k' or '-p' flag must be provided")
	}
}
