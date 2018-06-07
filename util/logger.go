package util

import (
	"os"

	"github.com/juju/loggo"
	"github.com/juju/loggo/loggocolor"
)

func getLogger(scope string) loggo.Logger {
	logger := loggo.GetLogger(scope)
	logger.SetLogLevel(loggo.DEBUG)
	loggo.ReplaceDefaultWriter(loggocolor.NewWriter(os.Stderr))

	return logger
}

func GetMainLogger() loggo.Logger {
	return getLogger("Main")
}

func GetMinerLogger(id string) loggo.Logger {
	return getLogger("Miner-" + id)
}

func GetBlockchainLogger() loggo.Logger {
	return getLogger("Blockchain")
}

func GetBlockLogger() loggo.Logger {
	return getLogger("Block")
}

func GetUserLogger(id string) loggo.Logger {
	return getLogger("User-" + id)
}

func GetBoosterLogger() loggo.Logger {
	return getLogger("Booster")
}

func GetTempLogger() loggo.Logger {
	return getLogger("Temp")
}
