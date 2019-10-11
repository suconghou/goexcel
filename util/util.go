package util

import (
	"log"
	"os"
)

// Log log to stdout
var (
	Log = log.New(os.Stdout, "", 0)
)
