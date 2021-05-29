package flusher_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFlusherTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FlusherTest Suite")
}
