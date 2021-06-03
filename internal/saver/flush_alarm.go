package saver

import (
	"fmt"
	"time"
)

type FlushAlarm interface {
	C() <-chan struct{}
	Close()
}

type flushAlarm struct {
	ticker  *time.Ticker
	ch      chan struct{}
	closeCh chan struct{}
}

func (this *flushAlarm) C() <-chan struct{} {
	return this.ch
}

func (this *flushAlarm) Close() {
	this.ticker.Stop()
	close(this.closeCh)
}

func NewFlushAlarmTicker(d time.Duration) (FlushAlarm, error) {
	if d <= 0 {
		return nil, fmt.Errorf("the duration must be greater zero, but got %v", d)
	}
	alarm := &flushAlarm{ticker: time.NewTicker(d), ch: make(chan struct{}), closeCh: make(chan struct{})}
	go func() {
		waitingLoop := true
		for waitingLoop {
			select {
			case v := <-alarm.ticker.C:
				_ = v
				select {
				case alarm.ch <- struct{}{}:
				default:
				}
			case v := <-alarm.closeCh:
				_ = v
				close(alarm.ch)
				waitingLoop = false
			}
		}
	}()

	return alarm, nil
}
