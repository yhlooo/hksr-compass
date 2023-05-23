package compass

import (
	"context"
	"testing"
	"time"

	"github.com/go-logr/logr"
)

// TestDefaultSolver 测试默认求解器
func TestDefaultSolver(t *testing.T) {
	// 创建一个求解器
	solver, err := NewDefaultSolver(SolverOptions{Logger: logr.Discard()})
	if err != nil {
		t.Errorf("new default solver error: %s", err)
		return
	}

	// 求解罗盘
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ret, err := solver.Solve(ctx, Compass{
		OuterRing:  Ring{Location: 0, Speed: 1},
		MiddleRing: Ring{Location: 4, Speed: -4},
		InnerRing:  Ring{Location: 0, Speed: 2},
		RingGroups: []RingGroup{
			OuterInnerRingGroup,
			OuterMiddleRingGroup,
			MiddleInnerRingGroup,
		},
	})
	if err != nil {
		t.Errorf("compass solve error: %s", err)
		return
	}

	// 校验结果
	expectedRet := Steps{{
		RingGroup: OuterInnerRingGroup,
		Count:     1,
	}, {
		RingGroup: OuterMiddleRingGroup,
		Count:     5,
	}, {
		RingGroup: MiddleInnerRingGroup,
		Count:     2,
	}}
	if ret.String() != expectedRet.String() {
		t.Errorf("unexpected result: %#v (expected: %#v)", ret.String(), expectedRet.String())
	}
}
