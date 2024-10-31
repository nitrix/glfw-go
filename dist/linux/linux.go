//go:build linux

package linux

// #cgo CFLAGS: -D_GLFW_X11 -O3
// #cgo LDFLAGS: -lm
import "C"
