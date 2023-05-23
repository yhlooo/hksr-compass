package compass

import (
	"context"

	"github.com/go-logr/logr"
)

// NewDefaultSolver 创建一个默认引航罗盘求解器
func NewDefaultSolver(opts SolverOptions) (Solver, error) {
	return &defaultSolver{
		logger: opts.Logger,
	}, nil
}

// defaultSolver 默认引航罗盘求解器
type defaultSolver struct {
	logger logr.Logger
}

var _ Solver = &defaultSolver{}

// Solve TODO 求解引航罗盘
func (s *defaultSolver) Solve(ctx context.Context, compass Compass) (Steps, error) {
	return nil, nil
}
