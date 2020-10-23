package integration_tests_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// TestIntegrationTests registers the integration test suite with ginkgo.
func TestIntegrationTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IntegrationTests Suite")
}

var _ = BeforeSuite(func() {
	// start up mongodb docker image
})

var _ = AfterSuite(func() {
	// shutdown mongodb docker image
})
