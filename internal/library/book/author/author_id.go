package author

import (
	"crypto/rand"
	"encoding/binary"

	"go.uber.org/zap"

	"github.com/MikhailKatarzhin/Library/pkg/logger"
	"github.com/MikhailKatarzhin/Library/pkg/process"
)

type ID uint64

func NewID() ID {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		logger.I().Error("can not to generate author id", zap.Error(err))
		process.Terminate()
	}

	return ID(binary.BigEndian.Uint64(b))
}
