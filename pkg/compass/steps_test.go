package compass

import (
	"testing"
)

// TestStepsString 测试 Steps.String 方法
func TestStepsString(t *testing.T) {
	steps := Steps{
		{RingGroup: InnerRingGroup, Count: 1},
		{RingGroup: OuterMiddleRingGroup, Count: 3},
		{RingGroup: MiddleInnerRingGroup, Count: 2},
		{RingGroup: OuterRingGroup, Count: 1},
		{RingGroup: InnerRingGroup, Count: 2},
		{RingGroup: MiddleRingGroup, Count: 0},
	}
	expectedRet := "i3,mi2,o1,om3"
	ret := steps.String()
	if ret != expectedRet {
		t.Errorf("unexpected result: %#v (expected: %#v)", ret, expectedRet)
	}
}
