package gbytes_test

import (
	. "github.com/botmetrics/go-botmetrics/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/botmetrics/go-botmetrics/Godeps/_workspace/src/github.com/onsi/gomega"

	"testing"
)

func TestGbytes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gbytes Suite")
}
