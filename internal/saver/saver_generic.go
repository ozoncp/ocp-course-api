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
	mu       *sync.Mutex
	blockCv  *sync.Cond
	closed   bool
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

	if this.closed {
		return
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
		for !this.closed && (len(this.buffer) >= this.capacity.ToInt()) {
			this.blockCv.Wait()
		}

		if !this.closed {
			this.buffer = append(this.buffer, v)
		}
	}
}

func (this *saverTValue) Close() {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.closed = true
	this.blockCv.Broadcast()
	this.alarm.Close()
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
	this.blockCv.Broadcast()
}

func (this *saverTValue) newBuffer() []TValue {
	return make([]TValue, 0, this.capacity.ToInt())
}

func (this *saverTValue) start() {
	go func() {
		for {
			select {
			case _, ok := <-this.alarm.C():
				if ok {
					this.flush()
				} else {
					return
				}
			case _, ok := <-this.closeCh:
				_ = ok
				return
			}
		}
	}()
}

func NewSaverTValue(
	flusher flusherTValue,
	alarm FlushAlarm,
	capacity commons.NaturalInt,
	os OverflowStrategy) SaverTValue {
	mu := &sync.Mutex{}
	res := &saverTValue{
		mu:       mu,
		blockCv:  sync.NewCond(mu),
		closed:   false,
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
