package compass

import (
	"testing"
)

// TestCompassString 测试 Compass.String 方法
func TestCompassString(t *testing.T) {
	ret := (&Compass{
		OuterRing:  Ring{Location: 0, Speed: 1},
		MiddleRing: Ring{Location: 4, Speed: -4},
		InnerRing:  Ring{Location: 0, Speed: 2},
		RingGroups: []RingGroup{
			OuterInnerRingGroup,
			OuterMiddleRingGroup,
			MiddleInnerRingGroup,
		},
	}).String()
	expectedRet := "0+1,4-4,0+2/mi,oi,om"

	if ret != expectedRet {
		t.Errorf("unexpected result: %#v (expected: %#v)", ret, expectedRet)
	}
}
