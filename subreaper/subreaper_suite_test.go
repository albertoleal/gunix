package subreaper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSubreaper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Subreaper Suite")
}
