package saver

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestFlusherAlarmClose(t *testing.T) {
	defer goleak.VerifyNone(t)
	alarm, err := NewFlushAlarmTicker(100 * time.Millisecond)
	assert.Equal(t, nil, err)
	_, ok := <-alarm.C()
	assert.Equal(t, true, ok)
	alarm.Close()
	for {
		_, ok := <-alarm.C()
		if !ok {
			break
		}
	}
}
