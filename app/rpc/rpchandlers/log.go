package rpchandlers

import (
	"github.com/Pugdag/pugdagd/infrastructure/logger"
	"github.com/Pugdag/pugdagd/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
