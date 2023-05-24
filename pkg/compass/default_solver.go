package compass

import (
	"context"
	"fmt"
	"sort"

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

// Solve 求解引航罗盘
func (s *defaultSolver) Solve(_ context.Context, compass Compass) (Steps, error) {
	// 校验入参
	if err := compass.Validate(); err != nil {
		return nil, fmt.Errorf("compass validation error: %w", err)
	}

	// 尝试所有可能的解法
	for _, solution := range s.getPossibleSolutions(compass) {
		if ok, _ := CheckSolution(compass, solution); ok {
			// 撞到了一个有效的解法
			return solution.Standardize(), nil
		}
	}

	return nil, fmt.Errorf("the compass has no solution")
}

// getPossibleSolutions 获取所有可能的解法
func (s *defaultSolver) getPossibleSolutions(compass Compass) []Steps {
	var possibleSolutions []Steps

	for _, rg := range compass.RingGroups {
		var temp []Steps
		// 因为转 6 次可以保证任何 rg 都能回到原点
		// TODO: 其实可以优化成使用 rg 涉及各圈循环周期的最小公约数
		for i := 0; i < 6; i++ {
			if len(possibleSolutions) == 0 {
				temp = append(temp, Steps{{
					RingGroup: rg,
					Count:     i,
				}})
				continue
			}
			for _, cur := range possibleSolutions {
				temp = append(temp, append(cur, Step{
					RingGroup: rg,
					Count:     i,
				}))
			}
		}
		possibleSolutions = temp
	}

	// 按步骤数排序
	sort.SliceStable(possibleSolutions, func(i, j int) bool {
		sumI := 0
		for _, step := range possibleSolutions[i] {
			sumI += step.Count
		}
		sumJ := 0
		for _, step := range possibleSolutions[j] {
			sumI += step.Count
		}
		return sumI < sumJ
	})

	return possibleSolutions
}
