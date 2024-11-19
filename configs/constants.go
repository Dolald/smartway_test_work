package configs

import "time"

const (
	ContextTime    = 2 * time.Second
	MaxHeaderBytes = 1 >> 20
	ReadTimeout    = 10 * time.Second
	WriteTimeout   = 10 * time.Second
	Id             = "id"
)
