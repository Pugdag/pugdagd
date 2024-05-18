package protocol

import (
	"github.com/Pugdag/pugdagd/infrastructure/logger"
	"github.com/Pugdag/pugdagd/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
