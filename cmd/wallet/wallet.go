package wallet

import (
	"github.com/urfave/cli"
)

var WalletCmd = cli.Command{
	Name:        "wallet",
	Aliases:     []string{"w"},
	Usage:       "options for wallet",
	Subcommands: []cli.Command{},
}

func Cmd() cli.Command {
	WalletCmd.Subcommands = append(WalletCmd.Subcommands, Create())

	return WalletCmd
}

func Create() cli.Command {
	return cli.Command{
		Name:  "create",
		Usage: "clink wallet create",
		Action: func(c *cli.Context) error {
			return create()
		},
	}
}

// TODO : implements me w/ test case
func create() error {
	return nil
}
