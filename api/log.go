package api

import (
	goutillog "github.com/spudtrooper/goutil/log"
)

func makeLog(prefix string) goutillog.Logger {
	return goutillog.MakeLog(prefix, goutillog.MakeLogColor(true))
}
