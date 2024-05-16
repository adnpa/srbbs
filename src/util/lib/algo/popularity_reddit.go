package algo

import (
	"math"
	"srbbs/src/util/lib"
	"time"
)

// reddit的排名算法

// Hot 帖子排名算法
func Hot(ups, downs int, date time.Time) float64 {
	// todo 后面改为上线时间
	ts := epochSeconds(time.Date(2005, 12, 8, 7, 46, 43, 0, time.UTC))
	x := ups - downs
	y := float64(lib.Sign(float64(x)))
	z := lib.Threshold(float64(x))
	return math.Round(math.Log10(z) + (y*ts)/45000)
}

// 帖子新旧程度
func epochSeconds(date time.Time) float64 {
	epoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	diff := date.Sub(epoch)
	return diff.Seconds()
}

// HotComment 评论排名算法
func HotComment(ups, downs float64) float64 {
	return _confidence(ups, downs)
}

func _confidence(ups, downs float64) float64 {
	n := ups + downs
	if n == 0 {
		return 0
	}

	z := 1.281551565545
	p := ups / n

	left := p + 1/(2*n)*z*z
	right := z * math.Sqrt(p*(1-p)/n+z*z/(4*n*n))
	under := 1 + 1/n*z*z

	return (left + right) / under
}
