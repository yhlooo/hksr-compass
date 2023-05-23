package compass

import (
	"context"

	"github.com/go-logr/logr"
)

// Solver 引航罗盘求解器
type Solver interface {
	// Solve 求解引航罗盘
	Solve(ctx context.Context, compass Compass) (Steps, error)
}

// SolverOptions Solver 的设置选项
type SolverOptions struct {
	// 日志记录器
	Logger logr.Logger
}
