//go:build darwin

package darwin

// #cgo CFLAGS: -D_GLFW_COCOA -O3
// #cgo LDFLAGS: -framework CoreFoundation -framework Cocoa -framework IOKit -framework QuartzCore
import "C"
