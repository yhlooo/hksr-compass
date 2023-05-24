package solve

import (
	"fmt"

	"github.com/bombsimon/logrusr/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/keybrl/hksr-compass/pkg/compass"
)

// Cmd solve 命令
var Cmd = &cobra.Command{
	Use:   "solve COMPASS_EXPRESSION",
	Short: "Solve a Navigation Compass.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logrusr.New(logrus.StandardLogger())
		// 创建求解器
		solver, err := compass.NewDefaultSolver(compass.SolverOptions{
			Logger: logger,
		})
		if err != nil {
			logger.Error(err, "new solver for navigation compass error")
			return fmt.Errorf("new solver for navigation compass error: %w", err)
		}
		// 解析输入罗盘
		input, err := compass.ParseCompass(args[0])
		if err != nil {
			logger.Error(err, "parse compass error")
			return fmt.Errorf("parse compass error: %w", err)
		}
		// 求解罗盘
		solution, err := solver.Solve(cmd.Context(), input)
		if err != nil {
			logger.Error(err, "solve navigation compass error")
			return fmt.Errorf("solve navigation compass error: %w", err)
		}
		fmt.Printf("Compass:  %s\n", input.String())
		fmt.Printf("Solution: %s\n", solution.String())
		return nil
	},
}
