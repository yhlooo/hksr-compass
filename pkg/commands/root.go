package commands

import (
	"github.com/spf13/cobra"

	"github.com/keybrl/hksr-compass/pkg/commands/solve"
)

const (
	binName = "hksr-compass"
	version = "0.1.0"
)

var (
	flagVerbose int
)

// Cmd 根命令
var Cmd = &cobra.Command{
	Use:     binName,
	Short:   "A tool for solving the Navigation Compass in the game Honkai: Star Rail.",
	Version: version,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	Cmd.PersistentFlags().CountVarP(&flagVerbose, "verbose", "v", "number for the log level verbosity")

	Cmd.AddCommand(
		solve.Cmd,
	)
}
