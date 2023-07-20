// filename: stuff.go
package main

// CGO_ENABLED=1 go build -buildmode=c-shared -o stuff.dll stuff.go
// cp stuff.dll stuff.h ../../

// #cgo CFLAGS: -g
import (
	"C"
)

//export Hello
func Hello() {
	println("hello world")
}

func main() {}
