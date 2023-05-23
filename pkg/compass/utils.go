package compass

import "fmt"

// CheckSolution 检查罗盘解决结果
func CheckSolution(compass Compass, steps Steps) (bool, error) {
	if err := compass.Validate(); err != nil {
		return false, fmt.Errorf("compass validation error: %w", err)
	}
	if err := steps.Validate(); err != nil {
		return false, fmt.Errorf("steps validation error: %w", err)
	}

	// 各圈初始位置
	inner := compass.InnerRing.Location
	middle := compass.MiddleRing.Location
	outer := compass.OuterRing.Location

	// 转一下
	for _, s := range steps {
		if !compass.IsRingGroupSupported(s.RingGroup) {
			return false, fmt.Errorf(
				"steps contains ring group not supported by compass: %s (must be one of %v)",
				s.RingGroup,
				compass.RingGroups,
			)
		}

		switch s.RingGroup {
		case InnerRingGroup:
			inner += s.Count * compass.InnerRing.Speed
		case MiddleRingGroup:
			middle += s.Count * compass.MiddleRing.Speed
		case MiddleInnerRingGroup:
			middle += s.Count * compass.MiddleRing.Speed
			inner += s.Count * compass.InnerRing.Speed
		case OuterRingGroup:
			outer += s.Count * compass.OuterRing.Speed
		case OuterMiddleRingGroup:
			outer += s.Count * compass.OuterRing.Speed
			middle += s.Count * compass.MiddleRing.Speed
		case OuterInnerRingGroup:
			outer += s.Count * compass.OuterRing.Speed
			inner += s.Count * compass.InnerRing.Speed
		}
	}

	// 检查各圈最终位置
	if inner%6 != 0 || middle%6 != 0 || outer%6 != 0 {
		return false, nil
	}
	return true, nil
}