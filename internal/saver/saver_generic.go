//go:generate genny -in=$GOFILE -out=$GOFILE.gen.go gen "TValue=model.Course,model.Lesson"
package saver

import (
	"sync"

	"github.com/cheekybits/genny/generic"

	"github.com/ozoncp/ocp-course-api/internal/utils/commons"
)

type TValue = generic.Type

type SaverTValue interface {
	SaveTValue(v TValue)
	Close()
}

type flusherTValue interface {
	FlushTValue(vs []TValue) []TValue
}

type saverTValue struct {
	mu       sync.Mutex
	blockCh  chan struct{}
	closeCh  chan struct{}
	capacity commons.NaturalInt
	os       OverflowStrategy
	buffer   []TValue
	flusher  flusherTValue
	alarm    FlushAlarm
}

func (this *saverTValue) SaveTValue(v TValue) {
	this.mu.Lock()
	defer this.mu.Unlock()

	if !this.os.new {
		panic("the instance of OverflowStrategy was wrong created")
	}

	switch this.os {
	case OverflowStrategyDropFirst():
		if len(this.buffer) < this.capacity.ToInt() {
			this.buffer = append(this.buffer, v)
		} else {
			this.buffer = append(this.buffer[1:], v)
		}

	case OverflowStrategyDropAll():
		if len(this.buffer) < this.capacity.ToInt() {
			this.buffer = append(this.buffer, v)
		} else {
			this.buffer = append(make([]TValue, 0, this.capacity.ToInt()), v)
		}

	case OverflowStrategyBlock():
		if len(this.buffer) < this.capacity.ToInt() {
			this.buffer = append(this.buffer, v)
		} else {
			this.mu.Unlock()
		waitingLoop:
			for {
				_, ok := <-this.blockCh
				this.mu.Lock()
				switch {
				case !ok:
					return
				case len(this.buffer) == this.capacity.ToInt():
					this.mu.Unlock()
					continue
				default:
					this.buffer = append(this.buffer, v)
					break waitingLoop
				}
			}
		}
	}
}

func (this *saverTValue) Close() {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.alarm.Close()
	close(this.blockCh)
	close(this.closeCh)
	if len(this.buffer) > 0 {
		this.buffer = this.flusher.FlushTValue(this.buffer)
	}
}

func (this *saverTValue) flush() {
	this.mu.Lock()
	defer this.mu.Unlock()
	if len(this.buffer) == 0 {
		return
	}
	this.buffer = this.flusher.FlushTValue(this.buffer)
	if cap(this.buffer) < this.capacity.ToInt() {
		this.buffer = append(this.newBuffer(), this.buffer...)
	}
	func() {
		defer func() {
			err := recover() //mute possible panic if the blockCh is closed
			_ = err
		}()
		select {
		case this.blockCh <- struct{}{}:
		default: //nobody are waiting data from the blockCh
		}
	}()
}

func (this *saverTValue) newBuffer() []TValue {
	return make([]TValue, 0, this.capacity.ToInt())
}

func (this *saverTValue) start() {
	go func() {
	waitingLoop:
		for {
			select {
			case _, ok := <-this.alarm.C():
				if ok {
					this.flush()
				} else {
					break waitingLoop
				}
			case _, ok := <-this.closeCh:
				_ = ok
				break waitingLoop
			}
		}
	}()
}

func NewSaverTValue(
	flusher flusherTValue,
	alarm FlushAlarm,
	capacity commons.NaturalInt,
	os OverflowStrategy) SaverTValue {
	res := &saverTValue{
		mu:       sync.Mutex{},
		blockCh:  make(chan struct{}),
		closeCh:  make(chan struct{}),
		os:       os,
		capacity: capacity,
		buffer:   make([]TValue, 0, capacity.ToInt()),
		flusher:  flusher,
		alarm:    alarm,
	}
	res.start()
	return res
}
