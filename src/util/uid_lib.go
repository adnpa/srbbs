package util

import (
	"srbbs/src/util/lib/algo"
)

func GetUidWithSnowFlake() (id uint64, err error) {
	return algo.GetID()
}
