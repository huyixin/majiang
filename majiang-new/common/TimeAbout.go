package common

import (
	"time"
)

type TimeCal struct {
	nanoTime int64
}

func NewTimeCal() *TimeCal {
	return &TimeCal{
		nanoTime: time.Now().UnixNano(),
	}
}

func (timer *TimeCal) ReStartCal() {
	timer.nanoTime = time.Now().UnixNano()
}

func (timer *TimeCal) CalCostTime() int64 {
	end := time.Now().UnixNano()
	return (end - timer.nanoTime) / 1000000 //10^6 (毫秒)
}
