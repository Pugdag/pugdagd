package standalone

import (
	"github.com/pugdag/pugdagd/infrastructure/logger"
	"github.com/pugdag/pugdagd/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)
