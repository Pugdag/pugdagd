package protocol

import (
	"github.com/pugdag/pugdagd/infrastructure/logger"
	"github.com/pugdag/pugdagd/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
