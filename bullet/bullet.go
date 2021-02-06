package main


// #cgo pkg-config: bullet
// #cgo windows LDFLAGS: -Wl,--allow-multiple-definition
// #include "c/SharedMemory/SharedMemoryInProcessPhysicsC_API.h"
import "C"

func main() {
	C.b3DisconnectSharedMemory(1)
}