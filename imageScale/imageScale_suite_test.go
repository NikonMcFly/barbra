package imageScale_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestImageScale(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ImageScale Suite")
}
