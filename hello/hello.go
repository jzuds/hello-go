package hello

import (
	"runtime"
)

func MyVersion() string {

	return runtime.Version()

}
