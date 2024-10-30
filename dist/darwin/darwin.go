//go:build darwin

package darwin

// #cgo CFLAGS: -D_GLFW_COCOA
// #cgo LDFLAGS: -framework CoreFoundation -framework Cocoa -framework IOKit -framework QuartzCore
import "C"
