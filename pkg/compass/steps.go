package compass

import (
	"fmt"
	"sort"
	"strings"
)

// Step 转动引航罗盘的步骤
type Step struct {
	// 转动的圈分组
	RingGroup RingGroup
	// 转动次数
	Count int
}

// ringGroupShortNames RingGroup 的缩写
var ringGroupShortNames = map[RingGroup]string{
	OuterRingGroup:       "o",
	MiddleRingGroup:      "m",
	InnerRingGroup:       "i",
	OuterMiddleRingGroup: "om",
	OuterInnerRingGroup:  "oi",
	MiddleInnerRingGroup: "mi",
}

// String 转为字符串表示
func (step *Step) String() string {
	if step == nil || step.Count <= 0 {
		return ""
	}
	return fmt.Sprintf("%s%d", ringGroupShortNames[step.RingGroup], step.Count)
}

// Validate 合法化
func (step *Step) Validate() error {
	return nil
}

// Steps 转动引航罗盘的步骤组合
type Steps []Step

// Standardize 标准化
func (steps Steps) Standardize() Steps {
	if len(steps) == 0 {
		return nil
	}
	// 拷贝一份出来处理
	sorted := make(Steps, len(steps))
	copy(sorted, steps)
	// 排序
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].RingGroup < sorted[j].RingGroup
	})
	// 合并
	var simplified Steps
	for _, s := range sorted {
		if s.Count <= 0 {
			continue
		}
		if len(simplified) > 0 && simplified[len(simplified)-1].RingGroup == s.RingGroup {
			simplified[len(simplified)-1].Count += s.Count
		} else {
			simplified = append(simplified, s)
		}
	}
	return simplified
}

// String 转为字符串表示
func (steps Steps) String() string {
	// 标准化
	std := steps.Standardize()
	if len(std) == 0 {
		return ""
	}
	// 逐个步骤转为字符串
	stepStrs := make([]string, len(std))
	for i := range std {
		stepStrs[i] = std[i].String()
	}
	// 逗号连接
	return strings.Join(stepStrs, ",")
}

// Validate TODO 合法化
func (steps Steps) Validate() error {
	return nil
}
