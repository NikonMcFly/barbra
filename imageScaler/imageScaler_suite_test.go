package imageScaler_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestImageScaler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ImageScaler Suite")
}
