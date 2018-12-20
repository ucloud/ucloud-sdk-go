package ucloud

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"
)

var (
	actionsLoggingLevels = make(map[string]log.Level)
)

// SetLogLevelByAction will set logging level by action name
func SetLogLevelByAction(name string, level log.Level) {
	actionsLoggingLevels[name] = level
}
