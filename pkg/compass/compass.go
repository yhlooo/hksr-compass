package compass

import (
	"context"

	"github.com/go-logr/logr"
)

// Ring 引航罗盘中的一圈
type Ring struct {
	// 位置
	// 按 12 小时制时钟的整点位置表示，比如 9 表示目标位置（即罗盘正左方）
	// 由于罗盘最小旋转单位为 60 度，因此该字段有效的值只有： 1, 3, 5, 7, 9, 11
	Location int
	// 旋转速度
	// 单位为 60 度，符号表示旋转方向，正数表示顺时针旋转，负数表示逆时针旋转
	// 比如： -1 表示每次逆时针旋转 60 度； 2 表示每次顺时针旋转 120 度
	Speed int
}

// RingGroup 引航罗盘圈分组
type RingGroup string

// RingGroup 的合法值
const (
	OuterRingGroup       RingGroup = "Outer"
	MiddleRingGroup      RingGroup = "Middle"
	InnerRingGroup       RingGroup = "Inner"
	OuterMiddleRingGroup RingGroup = "OuterMiddle"
	OuterInnerRingGroup  RingGroup = "OuterInner"
	MiddleInnerRingGroup RingGroup = "MiddleInner"
)

// Compass 引航罗盘
type Compass struct {
	// 内圈
	InnerRing Ring
	// 中圈
	MiddleRing Ring
	// 外圈
	OuterRing Ring
	// 圈分组
	// 可以同时旋转的一个或多个圈组成一个分组
	RingGroups []RingGroup
}

// Step 转动引航罗盘的步骤
type Step struct {
	// 转动的圈分组
	RingGroup RingGroup
	// 转动次数
	Count int
}

// Solver 引航罗盘求解器
type Solver interface {
	// Solve 求解引航罗盘
	Solve(ctx context.Context, compass Compass) ([]Step, error)
}

// SolverOptions Solver 的设置选项
type SolverOptions struct {
	// 日志记录器
	Logger logr.Logger
}
