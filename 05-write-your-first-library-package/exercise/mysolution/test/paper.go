package test

import (
	"runtime"
)

func Paper() string{
	// return "HELLO WORLD"
	return runtime.Version()
}
