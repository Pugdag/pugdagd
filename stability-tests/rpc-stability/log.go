package main

import (
	"github.com/Pugdag/pugdagd/infrastructure/logger"
	"github.com/Pugdag/pugdagd/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("JSTT")
	spawn      = panics.GoroutineWrapperFunc(log)
)
