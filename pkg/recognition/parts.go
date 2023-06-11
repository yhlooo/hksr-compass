package recognition

import (
	"embed"
	"fmt"

	"github.com/keybrl/hksr-compass/pkg/compass"
)

// partImages 罗盘的局部图片
//
//go:embed assests/*.png
var partImages embed.FS

// getPart 获取罗盘局部图片
func getPart(ring compass.RingGroup, speed int) []byte {
	data, err := partImages.ReadFile(getPartName(ring, speed))
	if err != nil {
		// TODO: 理论上应该打个日志
		data = nil
	}
	return data
}

// getPartName 获取罗盘局部图片名
func getPartName(ring compass.RingGroup, speed int) string {
	return fmt.Sprintf("assests/%s%+d.png", ring.ShortName(), speed)
}
