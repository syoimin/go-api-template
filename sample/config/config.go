package config

import (
	"os"
)

var (
	ENVIRONMENT  = os.Getenv("ENVIRONMENT")
	ALLOW_ORIGIN = os.Getenv("ALLOW_ORIGIN")
)
