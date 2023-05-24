package compass

import (
	"testing"
)

// TestParseCompass 测试 ParseCompass
func TestParseCompass(t *testing.T) {
	compass, err := ParseCompass("0+1,4-4,0+2/oi,om,mi")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}

	expectedRet := Compass{
		OuterRing:  Ring{Location: 0, Speed: 1},
		MiddleRing: Ring{Location: 4, Speed: -4},
		InnerRing:  Ring{Location: 0, Speed: 2},
		RingGroups: []RingGroup{
			OuterInnerRingGroup,
			OuterMiddleRingGroup,
			MiddleInnerRingGroup,
		},
	}

	if compass.String() != expectedRet.String() {
		t.Errorf("unexpected result: %#v (expected: %#v)", compass.String(), expectedRet.String())
	}
}
