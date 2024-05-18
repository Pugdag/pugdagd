package rpc

import (
	"github.com/pugdag/pugdagd/infrastructure/logger"
	"github.com/pugdag/pugdagd/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
