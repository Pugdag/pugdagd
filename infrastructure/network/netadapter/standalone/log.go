package standalone

import (
	"github.com/Pugdag/pugdagd/infrastructure/logger"
	"github.com/Pugdag/pugdagd/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)
