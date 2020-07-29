package utills

import "time"

const ServerUrl = "0.0.0.0:8080"

type timeouts struct {
	WriteTimeout   time.Duration
	ReadTimeout    time.Duration
	ContextTimeout time.Duration
}

var Timeouts timeouts
