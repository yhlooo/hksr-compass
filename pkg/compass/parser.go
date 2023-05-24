package compass

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	compassRegexpStr = `(?P<outerRing>[0-9-+]+),` +
		`(?P<middleRing>[0-9-+]+),` +
		`(?P<innerRing>[0-9-+]+)/` +
		`(?P<ringGroups>[imo,]+)`
	ringRegexpStr = `(?P<location>[0-5])(?P<speed>(?:\+|-)[1-4])`
)

var (
	compassRegexp = regexp.MustCompile(compassRegexpStr)
	ringRegexp    = regexp.MustCompile(ringRegexpStr)
)

// ParseCompass 解析字符串表示的罗盘信息
func ParseCompass(compass string) (Compass, error) {
	ret := Compass{}

	// 正则
	groups := compassRegexp.FindStringSubmatch(compass)
	if groups == nil {
		return ret, fmt.Errorf("invalid compass expression: \"%s\" (not match \"%s\")", compass, compassRegexpStr)
	}

	// 各捕获组分别解析

	outer, err := ParseRing(groups[compassRegexp.SubexpIndex("outerRing")])
	if err != nil {
		return ret, fmt.Errorf("parse outer ring error: %w", err)
	}
	ret.OuterRing = outer

	middle, err := ParseRing(groups[compassRegexp.SubexpIndex("middleRing")])
	if err != nil {
		return ret, fmt.Errorf("parse middle ring error: %w", err)
	}
	ret.MiddleRing = middle

	inner, err := ParseRing(groups[compassRegexp.SubexpIndex("innerRing")])
	if err != nil {
		return ret, fmt.Errorf("parse inner ring error: %w", err)
	}
	ret.InnerRing = inner

	rgs, err := ParseRingGroups(groups[compassRegexp.SubexpIndex("ringGroups")])
	if err != nil {
		return ret, fmt.Errorf("parse ring groups error: %w", err)
	}
	ret.RingGroups = rgs

	return ret, nil
}

// ParseRingGroups 解析字符串表示的罗盘圈组列表
func ParseRingGroups(ringGroups string) ([]RingGroup, error) {
	var ret []RingGroup
	// 按 , 切分各圈组解析
	for i, rgStr := range strings.Split(ringGroups, ",") {
		rg, err := ParseRingGroup(rgStr)
		if err != nil {
			return nil, fmt.Errorf("parse the ring group at index %d error: %w", i, err)
		}
		ret = append(ret, rg)
	}
	return ret, nil
}

// ParseRingGroup 解析字符串表示的罗盘圈组
func ParseRingGroup(ringGroup string) (RingGroup, error) {
	switch ringGroup {
	case "o":
		return OuterRingGroup, nil
	case "m":
		return MiddleRingGroup, nil
	case "i":
		return InnerRingGroup, nil
	case "oi", "io":
		return OuterInnerRingGroup, nil
	case "om", "mo":
		return OuterMiddleRingGroup, nil
	case "mi", "im":
		return MiddleInnerRingGroup, nil
	}
	return 0, fmt.Errorf("unknown ring group: %s", ringGroup)
}

// ParseRing 解析字符串表示的罗盘圈
func ParseRing(ring string) (Ring, error) {
	ret := Ring{}

	// 正则
	groups := ringRegexp.FindStringSubmatch(ring)
	if groups == nil {
		return ret, fmt.Errorf("invalid ring expression: \"%s\" (not match \"%s\")", ring, ringRegexpStr)
	}

	locationStr := groups[ringRegexp.SubexpIndex("location")]
	location, err := strconv.ParseInt(locationStr, 10, 8)
	if err != nil {
		return ret, fmt.Errorf("parse ring location \"%s\" error: %w", locationStr, err)
	}
	ret.Location = int(location)

	speedStr := groups[ringRegexp.SubexpIndex("speed")]
	speed, err := strconv.ParseInt(speedStr, 10, 8)
	if err != nil {
		return ret, fmt.Errorf("parse ring speed \"%s\" error: %w", speedStr, err)
	}
	ret.Speed = int(speed)

	return ret, nil
}
