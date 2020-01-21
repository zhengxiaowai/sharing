package sharer_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSharer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sharer Suite")
}
