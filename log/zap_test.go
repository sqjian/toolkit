package log_test

import (
	"github.com/sqjian/toolkit/log"
	"testing"
)

func TestGenZapLog(t *testing.T) {
	logger, _ := log.GenZapLog(
		"test.log",
		1,
		1,
		1,
		"debug",
		true,
	)

	logger.Info("testing...")
}
