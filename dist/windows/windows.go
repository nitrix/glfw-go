//go:build windows

package windows

// #cgo CFLAGS: -I../common -D_GLFW_WIN32 -O3
// #cgo LDFLAGS: -limm32 -luser32 -lkernel32 -lgdi32 -lshell32
import "C"
