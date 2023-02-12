package timer

import (
	"sync"
	"time"

	"github.com/cheggaaa/pb"
)

type Timer struct {
	Wg    *sync.WaitGroup
	Count int
	Over  bool
	Bar   *pb.ProgressBar
}

func NewTimer() *Timer {
	var wg sync.WaitGroup
	count := 100
	return &Timer{
		Wg:    &wg,
		Count: count,
		Bar:   pb.StartNew(count),
		Over:  false,
	}
}

func (timer *Timer) Start() {
	timer.Wg.Add(1)
	go func() {
		ticker := time.NewTicker(time.Millisecond * 15)
		for {
			<-ticker.C
			if timer.Bar.Get() < 99 {
				timer.Bar.Increment()
			}
			if timer.Over {
				for i := timer.Bar.Get(); i < 100; i++ {
					timer.Bar.Increment()
				}
				timer.Wg.Done()
				return
			}
		}
	}()
}

func (timer *Timer) Stop() {
	timer.Over = true
	timer.Wg.Wait()
	timer.Bar.Finish()
}
