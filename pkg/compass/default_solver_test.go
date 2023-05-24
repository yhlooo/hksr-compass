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
	expectedRets := []string{"mi2,oi4,om2", "mi2,oi1,om5"}
	correct := false
	for _, expectedRet := range expectedRets {
		if ret.String() == expectedRet {
			correct = true
			break
		}
	}
	if !correct {
		t.Errorf("unexpected result: %#v (expect to be one of %#v)", ret.String(), expectedRets)
	}
}
