package resource_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStaticResource(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StaticResource Suite")
}
