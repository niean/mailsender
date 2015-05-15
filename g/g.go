package g

import (
	"log"
	"runtime"
)

// changelog:
// 0.0.1: basic mail-sending service, http-api, readme. send mail concurrently
const (
	VERSION = "0.0.1"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
