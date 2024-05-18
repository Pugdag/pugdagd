package integration

import (
	"os"
	"testing"

	"github.com/Pugdag/pugdagd/infrastructure/logger"
)

func TestMain(m *testing.M) {
	logger.SetLogLevels(logger.LevelDebug)
	logger.InitLogStdout(logger.LevelDebug)

	os.Exit(m.Run())
}
