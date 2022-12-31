package process_test

import (
	"os"
	"syscall"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"gitlab.wildberries.ru/wb-branches/wb-branches/backend/wb-branches/pkg/process"
)

func TestWaitForTermination(t *testing.T) {
	t.Parallel()

	RegisterTestingT(t)

	go func() {
		time.Sleep(time.Millisecond)
		callSignal(syscall.SIGTERM)
	}()

	// If the test succeeded, the SIGTERM signal was intercepted.

	process.WaitForTermination()
}

func callSignal(signal os.Signal) {
	proc, err := os.FindProcess(os.Getpid())
	Ω(err).ShouldNot(HaveOccurred())

	err = proc.Signal(signal)
	Ω(err).ShouldNot(HaveOccurred())
}
