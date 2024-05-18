package main

import (
	"github.com/pugdag/pugdagd/infrastructure/logger"
	"github.com/pugdag/pugdagd/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("JSTT")
	spawn      = panics.GoroutineWrapperFunc(log)
)
