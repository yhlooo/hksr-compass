package solve

import (
	"fmt"

	"github.com/bombsimon/logrusr/v4"
	"github.com/keybrl/hksr-compass/pkg/compass"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Cmd solve 命令
var Cmd = &cobra.Command{
	Use:   "solve",
	Short: "Solve a Navigation Compass.",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logrusr.New(logrus.StandardLogger())
		solver, err := compass.NewDefaultSolver(compass.SolverOptions{
			Logger: logger,
		})
		if err != nil {
			logger.Error(err, "new solver for navigation compass error")
			return fmt.Errorf("new solver for navigation compass error: %w", err)
		}
		steps, err := solver.Solve(cmd.Context(), compass.Compass{
			// TODO: ...
		})
		if err != nil {
			logger.Error(err, "solve navigation compass error")
			return fmt.Errorf("solve navigation compass error: %w", err)
		}
		// TODO: ...
		_ = steps
		return nil
	},
}
