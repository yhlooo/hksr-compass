package recognition

import (
	"fmt"
	"image/color"

	"gocv.io/x/gocv"

	"github.com/keybrl/hksr-compass/pkg/compass"
)

// allRings 罗盘所有圈
var allRings = []compass.RingGroup{
	compass.OuterRingGroup,
	// compass.MiddleRingGroup,
	// compass.InnerRingGroup,
}

// RecognizeCompassImage TODO 识别罗盘图像
func RecognizeCompassImage(inputImage []byte) (*compass.Compass, error) {
	// 加载输入图片
	input, err := gocv.IMDecode(inputImage, gocv.IMReadGrayScale)
	if err != nil {
		return nil, fmt.Errorf("load image error: %w", err)
	}
	defer func() { _ = input.Close() }()

	// 逐个模版匹配
	part := gocv.NewMat()
	defer func() { _ = part.Close() }()
	for _, ring := range allRings {
		for speed := -4; speed <= 4; speed++ {
			if speed == 0 {
				continue
			}
			// 加载局部图片
			if err := gocv.IMDecodeIntoMat(getPart(ring, speed), gocv.IMReadGrayScale, &part); err != nil {
				return nil, fmt.Errorf("load compass part image error: %w", err)
			}
			// 查找罗盘局部
			matchCnt, score, location := matchCompassPart(input, part)
			fmt.Printf(
				"%s matchCnt: %d, score: %.3f, location: %d\n",
				getPartName(ring, speed),
				matchCnt,
				score,
				location,
			)
		}
	}
	return nil, nil
}

func matchCompassPart(img gocv.Mat, part gocv.Mat) (matchCnt int, score float64, location int) {
	window := gocv.NewWindow("Debug")
	defer func() { _ = window.Close() }()
	debugImage := gocv.NewMat()
	defer func() { _ = debugImage.Close() }()

	// 预处理（二值化）
	gocv.Threshold(img, &img, 150, 255, gocv.ThresholdBinary)
	gocv.Threshold(part, &part, 150, 255, gocv.ThresholdBinary)

	// 创建一个特征检测器（检测图片中角点、边缘等特征）
	featureDetector := gocv.NewSIFT()
	defer func() { _ = featureDetector.Close() }()

	mask := gocv.NewMat()
	defer func() { _ = mask.Close() }()

	// 特征检测
	imgKP, imgFeatures := featureDetector.DetectAndCompute(img, mask)
	partKP, partFeatures := featureDetector.DetectAndCompute(part, mask)

	// 创建一个匹配器
	matcher := gocv.NewBFMatcher()
	defer func() { _ = matcher.Close() }()

	// 特征匹配
	matchRet := matcher.KnnMatch(partFeatures, imgFeatures, 2)
	var matches []gocv.DMatch
	for _, m := range matchRet {
		// if len(m) > 0 {
		// 	matches = append(matches, m[0])
		// 	matchCnt++
		// 	score += m[0].Distance
		// }
		if m[0].Distance < m[1].Distance*0.75 {
			matches = append(matches, m[0])
			matchCnt++
			score += m[0].Distance
		}
	}

	gocv.DrawMatches(
		part, partKP,
		img, imgKP,
		matches, &debugImage,
		color.RGBA{255, 255, 255, 1}, color.RGBA{255, 255, 255, 1}, nil, gocv.DrawDefault,
	)
	window.IMShow(debugImage)
	window.WaitKey(2000)

	return
}
