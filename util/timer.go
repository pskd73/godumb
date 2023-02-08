package util

import (
	"fmt"
	"time"
)

type Timer struct {
	Checkpoints []time.Time
}

func (self *Timer) Init() {
	self.Checkpoints = []time.Time{}
}

func (self *Timer) Check(log bool) {
	self.Checkpoints = append(self.Checkpoints, time.Now())
	if log {
		self.LogLastTat()
	}
}

func (self *Timer) LogLastTat() {
	if len(self.Checkpoints) > 1 {
		lastTwo := self.Checkpoints[len(self.Checkpoints)-2:]
		fmt.Println(lastTwo[1].Sub(lastTwo[0]))
	}
}
