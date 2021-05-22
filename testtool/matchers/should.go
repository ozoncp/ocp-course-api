package matchers

import (
	"fmt"
	"reflect"
	"testing"
)

func shouldBe(result interface{}, expect interface{}) (string, bool) {
	if !reflect.DeepEqual(result, expect) {
		return fmt.Sprintf("%v doesn't equal %v.", result, expect), false
	}
	return "", true
}

// if the result equals the expect, execution continues. Otherwise, it fails test with Fatal.
func ShouldBe(t *testing.T, result interface{}, expect interface{}) {
	t.Helper()
	msg, res := shouldBe(result, expect)
	if !res {
		t.Fatal(msg)
	}
}

// if the result equals the expect, execution continues. Otherwise, it fails test with Fatal.
// Appends an error log with the clue.
func ShouldBe_with(t *testing.T, result interface{}, expect interface{}, clue string) {
	t.Helper()
	msg, res := shouldBe(result, expect)
	if !res {
		t.Fatal(msg + clue)
	}
}

// if the result equals the expect, execution continues. Otherwise, it fails test with Error.
func ShouldBeE(t *testing.T, result interface{}, expect interface{}) {
	t.Helper()
	msg, res := shouldBe(result, expect)
	if !res {
		t.Error(msg)
	}
}

// if the result equals the expect, execution continues. Otherwise, it fails test with Error.
// Appends an error log with the clue.
func ShouldBeE_with(t *testing.T, result interface{}, expect interface{}, clue string) {
	t.Helper()
	msg, res := shouldBe(result, expect)
	if !res {
		t.Error(msg + clue)
	}
}

func ShouldPanic(t *testing.T, code func()) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Expects panic, but the code didn't produce it.")
		}
	}()

	code()
}
