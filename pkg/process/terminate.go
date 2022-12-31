package process

import (
	"os"
	"syscall"

	"go.uber.org/zap"

	"gitlab.wildberries.ru/wb-branches/wb-branches/backend/wb-branches/pkg/logger"
)

// Terminate its application gratefully.
func Terminate() {
	pid := os.Getpid()

	proc, err := os.FindProcess(pid)
	if err != nil {
		logger.MustGetLogger().Fatal(
			"failed to find process ID for terminate myself",
			zap.Error(err),
		)
	}

	if err = proc.Signal(syscall.SIGTERM); err != nil {
		logger.MustGetLogger().Fatal(
			"failed to send SIGTERM signal to myself",
			zap.Int("PID", pid),
			zap.Error(err),
		)
	}
}
