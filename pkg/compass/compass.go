package compass

// Ring 引航罗盘中的一圈
type Ring struct {
	// 位置
	// 指针从目标位置（即罗盘正左方向）沿顺时针方向旋转到当前位置所需旋转的角度处以 60 度，
	// 比如 0 表示目标位置， 3 表示指针指向正右方向
	// 因为一周是 360 度，因此该字段有效范围是： 0-5
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

// Validate TODO 合法化
func (compass *Compass) Validate() error {
	return nil
}

// IsRingGroupSupported 判断指定圈分组是否是当前罗盘支持的
func (compass *Compass) IsRingGroupSupported(rg RingGroup) bool {
	if compass == nil {
		return false
	}
	for _, supportedRG := range compass.RingGroups {
		if supportedRG == rg {
			return true
		}
	}
	return false
}
