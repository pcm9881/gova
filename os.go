package gova

import (
	"fmt"
	"runtime"
)

func GetOS() string {
	goos := runtime.GOOS
	fmt.Printf("%s\n", goos)
	return goos
}
