package saver

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"

	"github.com/ozoncp/ocp-course-api/internal/mock_flusher"
	"github.com/ozoncp/ocp-course-api/internal/mock_saver"
	"github.com/ozoncp/ocp-course-api/internal/utils/commons"
)

func makeMocks(t *testing.T) (*gomock.Controller, *mock_flusher.MockFlusherTValue, *mock_saver.MockFlushAlarm) {
	ctrl := gomock.NewController(t)
	return ctrl, mock_flusher.NewMockFlusherTValue(ctrl), mock_saver.NewMockFlushAlarm(ctrl)
}

func TestSaverGenericClose(t *testing.T) {
	ctrl, mf, ma := makeMocks(t)
	defer ctrl.Finish()

	alarmCh := make(chan struct{})
	defer close(alarmCh)

	func() {
		defer goleak.VerifyNone(t)

		ma.EXPECT().C().Return(alarmCh).MinTimes(1)
		ma.EXPECT().Close().Times(1)
		mf.EXPECT().FlushTValue([]TValue{33}).Return([]TValue{}).Times(1)

		s := NewSaverTValue(
			mf,
			ma,
			commons.NewNaturalIntPanic(1),
			OverflowStrategyDropFirst(),
		)

		time.Sleep(100 * time.Millisecond)

		s.SaveTValue(33)
		s.Close()

		time.Sleep(100 * time.Millisecond)
	}()
}

func TestSaverGenericFlushByAlarm(t *testing.T) {
	defer goleak.VerifyNone(t)
	ctrl, mf, ma := makeMocks(t)
	defer ctrl.Finish()

	alarmCh := make(chan struct{})
	defer close(alarmCh)

	ma.EXPECT().C().Return(alarmCh).MinTimes(1)
	ma.EXPECT().Close().Times(1)
	s := NewSaverTValue(
		mf,
		ma,
		commons.NewNaturalIntPanic(1),
		OverflowStrategyDropFirst(),
	)

	mf.EXPECT().FlushTValue([]TValue{33}).Return([]TValue{}).Times(1)
	mf.EXPECT().FlushTValue([]TValue{42}).Return([]TValue{}).Times(1)

	time.Sleep(100 * time.Millisecond)
	s.SaveTValue(33)

	alarmCh <- struct{}{}
	time.Sleep(100 * time.Millisecond)

	s.SaveTValue(42)
	s.Close()
	time.Sleep(100 * time.Millisecond)
}

func TestSaverGenericCheckOverflowStrategyDropFirst(t *testing.T) {
	defer goleak.VerifyNone(t)
	ctrl, mf, ma := makeMocks(t)
	defer ctrl.Finish()

	alarmCh := make(chan struct{})
	defer close(alarmCh)

	ma.EXPECT().C().Return(alarmCh).MinTimes(1)
	ma.EXPECT().Close().Times(1)
	s := NewSaverTValue(
		mf,
		ma,
		commons.NewNaturalIntPanic(2),
		OverflowStrategyDropFirst(),
	)

	mf.EXPECT().FlushTValue([]TValue{2, 3}).Return([]TValue{}).Times(1)

	time.Sleep(100 * time.Millisecond)

	s.SaveTValue(1)
	s.SaveTValue(2)
	s.SaveTValue(3)

	alarmCh <- struct{}{}
	time.Sleep(100 * time.Millisecond)

	s.Close()
}

func TestSaverGenericCheckOverflowStrategyDropAll(t *testing.T) {
	defer goleak.VerifyNone(t)
	ctrl, mf, ma := makeMocks(t)
	defer ctrl.Finish()

	alarmCh := make(chan struct{})
	defer close(alarmCh)

	ma.EXPECT().C().Return(alarmCh).MinTimes(1)
	ma.EXPECT().Close().Times(1)
	s := NewSaverTValue(
		mf,
		ma,
		commons.NewNaturalIntPanic(2),
		OverflowStrategyDropAll(),
	)

	mf.EXPECT().FlushTValue([]TValue{3}).Return([]TValue{}).Times(1)

	time.Sleep(100 * time.Millisecond)

	s.SaveTValue(1)
	s.SaveTValue(2)
	s.SaveTValue(3)

	alarmCh <- struct{}{}
	time.Sleep(100 * time.Millisecond)

	s.Close()
}

func TestSaverGenericCheckOverflowStrategyBlock(t *testing.T) {
	defer goleak.VerifyNone(t)
	ctrl, mf, ma := makeMocks(t)
	defer ctrl.Finish()

	alarmCh := make(chan struct{})
	defer close(alarmCh)

	ma.EXPECT().C().Return(alarmCh).MinTimes(1)
	ma.EXPECT().Close().Times(1)
	s := NewSaverTValue(
		mf,
		ma,
		commons.NewNaturalIntPanic(2),
		OverflowStrategyBlock(),
	)

	mf.EXPECT().FlushTValue([]TValue{1, 2}).Return([]TValue{}).Times(1)
	mf.EXPECT().FlushTValue([]TValue{3}).Return([]TValue{}).Times(1)

	time.Sleep(100 * time.Millisecond)

	s.SaveTValue(1)
	s.SaveTValue(2)

	ch := make(chan struct{})
	defer close(ch)
	go func() {
		s.SaveTValue(3)
		ch <- struct{}{}
	}()

	time.Sleep(100 * time.Millisecond)
	select {
	case <-ch:
		assert.Fail(t, "SaveTValue is not blocking")
	default:
	}

	alarmCh <- struct{}{}
	time.Sleep(100 * time.Millisecond)

	select {
	case <-ch:
	default:
		assert.Fail(t, "SaveTValue is still blocking")
	}

	s.Close()
	time.Sleep(100 * time.Millisecond)
}
