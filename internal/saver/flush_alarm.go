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
		for {
			select {
			case <-alarm.ticker.C:
				select {
				case alarm.ch <- struct{}{}:
				default:
				}
			case <-alarm.closeCh:
				close(alarm.ch)
				return
			}
		}
	}()

	return alarm, nil
}
