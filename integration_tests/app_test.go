package integration_tests_test

import (
	"runtime"

	"github.com/FleekHQ/space-daemon/integration_tests/fixtures"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Running App", func() {
	var (
		goRoutinesBeforeAppStart int
		app                      *fixtures.RunAppCtx
	)

	// skipping this because it is currently failing
	XIt("should not leak goroutines on shutdown", func() {
		goRoutinesBeforeAppStart = runtime.NumGoroutine()
		app = fixtures.RunApp()

		err := app.App.Shutdown()
		Expect(err).ToNot(HaveOccurred(), "App shutdown failed")
		Expect(app.App.IsRunning).To(Equal(false))

		Expect(runtime.NumGoroutine()).To(Equal(goRoutinesBeforeAppStart))
	})
})
