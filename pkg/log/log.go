package log

import (
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

// MLog is a base parent logger for the microservice
var MLog = logf.Log.WithName("template")
