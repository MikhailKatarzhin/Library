// Package id implements id that use in book and its entity.
package id

import (
	"crypto/rand"
	"encoding/binary"
	"go.uber.org/zap"

	"github.com/MikhailKatarzhin/Library/pkg/logger"
	"github.com/MikhailKatarzhin/Library/pkg/process"
)

const int64ByteCount = 8

type ID uint64

func NewID() ID {
	b := make([]byte, int64ByteCount)
	if _, err := rand.Read(b); err != nil {
		logger.I().Error("can not to generate id", zap.Error(err))
		process.Terminate()
	}

	return ID(binary.BigEndian.Uint64(b))
}
