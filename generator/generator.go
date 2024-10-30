package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
)

func copyFile(src, dst string) {
	_ = os.MkdirAll(filepath.Dir(dst), 0750)
	srcFile, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	defer dstFile.Close()

	_, err = srcFile.WriteTo(dstFile)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Copied file %s to %s\n", src, dst)
}

func extractGlfwConstants() []string {
	out := []string{}

	content, err := os.ReadFile("thirdparty/glfw/include/GLFW/glfw3.h")
	if err != nil {
		panic(err)
	}

	renameConstant := func(name string) string {
		name = strcase.ToCamel(name)
		name = strings.ReplaceAll(name, "Opengl", "OpenGL")
		if name == "OpenGLForwardCompat" {
			return "OpenGLForwardCompatible"
		}
		return name
	}

	knownConstants := map[string]string{}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "#define GLFW_") {
			line = strings.TrimSpace(line)
			line = strings.ReplaceAll(line, "\t", " ")
			for {
				before := line
				line = strings.ReplaceAll(line, "  ", " ")
				if before == line {
					break
				}
			}

			parts := strings.Split(line, " ")
			_, err := strconv.Atoi(parts[2])
			if err == nil || strings.HasPrefix(parts[2], "0x") {
				knownConstants[parts[1]] = parts[2]
				// out = append(out, fmt.Sprintf("const %s = %s", renameConstant(strings.TrimPrefix(parts[1], "GLFW_")), parts[2]))
				out = append(out, fmt.Sprintf("const %s = %s", renameConstant(strings.TrimPrefix(parts[1], "GLFW_")), "C."+parts[1]))
			} else if knownConstants[parts[2]] != "" {
				// out = append(out, fmt.Sprintf("const %s = %s", renameConstant(strings.TrimPrefix(parts[1], "GLFW_")), knownConstants[parts[2]]))
				out = append(out, fmt.Sprintf("const %s = %s", renameConstant(strings.TrimPrefix(parts[1], "GLFW_")), "C."+parts[1]))
			}
		}
	}

	return out
}

func generateGlfwConstants() {
	preamble := "package glfw\n"
	preamble += "\n"
	preamble += "// #cgo CFLAGS: -Idist/include\n"
	preamble += "// #include \"GLFW/glfw3.h\"\n"
	preamble += "import \"C\"\n"
	preamble += "\n"

	glfwConstants := extractGlfwConstants()
	err := os.WriteFile("constants.go", []byte(preamble+strings.Join(glfwConstants, "\n")), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Copied %d constants to constants.go\n", len(glfwConstants))
}

func main() {
	// "thirdparty/glfw/include/GLFW/glfw3.h",
	// "thirdparty/glfw/include/GLFW/glfw3native.h",

	commonFiles := []string{
		"thirdparty/glfw/src/internal.h",
		"thirdparty/glfw/src/mappings.h",
		"thirdparty/glfw/src/platform.h",
		"thirdparty/glfw/src/context.c",
		"thirdparty/glfw/src/init.c",
		"thirdparty/glfw/src/input.c",
		"thirdparty/glfw/src/monitor.c",
		"thirdparty/glfw/src/platform.c",
		"thirdparty/glfw/src/vulkan.c",
		"thirdparty/glfw/src/window.c",
		"thirdparty/glfw/src/egl_context.c",
		"thirdparty/glfw/src/osmesa_context.c",
		"thirdparty/glfw/src/null_init.c",
		"thirdparty/glfw/src/null_monitor.c",
		"thirdparty/glfw/src/null_window.c",
		"thirdparty/glfw/src/null_joystick.c",
		"thirdparty/glfw/src/null_platform.h",
		"thirdparty/glfw/src/null_joystick.h",
	}

	// Darwin aka Mac OS (cocoa)
	for _, file := range commonFiles {
		copyFile(file, "dist/darwin/"+filepath.Base(file))
	}
	copyFile("thirdparty/glfw/src/cocoa_time.h", "dist/darwin/cocoa_time.h")
	copyFile("thirdparty/glfw/src/posix_thread.h", "dist/darwin/posix_thread.h")
	copyFile("thirdparty/glfw/src/cocoa_platform.h", "dist/darwin/cocoa_platform.h")
	copyFile("thirdparty/glfw/src/cocoa_joystick.h", "dist/darwin/cocoa_joystick.h")
	copyFile("thirdparty/glfw/src/cocoa_init.m", "dist/darwin/cocoa_init.m")
	copyFile("thirdparty/glfw/src/cocoa_joystick.m", "dist/darwin/cocoa_joystick.m")
	copyFile("thirdparty/glfw/src/cocoa_monitor.m", "dist/darwin/cocoa_monitor.m")
	copyFile("thirdparty/glfw/src/cocoa_window.m", "dist/darwin/cocoa_window.m")
	copyFile("thirdparty/glfw/src/nsgl_context.m", "dist/darwin/nsgl_context.m")
	copyFile("thirdparty/glfw/src/cocoa_time.c", "dist/darwin/cocoa_time.c")
	copyFile("thirdparty/glfw/src/posix_module.c", "dist/darwin/posix_module.c")
	copyFile("thirdparty/glfw/src/posix_thread.c", "dist/darwin/posix_thread.c")

	// Windows
	for _, file := range commonFiles {
		copyFile(file, "dist/windows/"+filepath.Base(file))
	}
	copyFile("thirdparty/glfw/src/win32_time.h", "dist/windows/win32_time.h")
	copyFile("thirdparty/glfw/src/win32_thread.h", "dist/windows/win32_thread.h")
	copyFile("thirdparty/glfw/src/win32_platform.h", "dist/windows/win32_platform.h")
	copyFile("thirdparty/glfw/src/win32_joystick.h", "dist/windows/win32_joystick.h")
	copyFile("thirdparty/glfw/src/win32_thread.c", "dist/windows/win32_thread.c")
	copyFile("thirdparty/glfw/src/win32_time.c", "dist/windows/win32_time.c")
	copyFile("thirdparty/glfw/src/win32_module.c", "dist/windows/win32_module.c")
	copyFile("thirdparty/glfw/src/win32_init.c", "dist/windows/win32_init.c")
	copyFile("thirdparty/glfw/src/win32_joystick.c", "dist/windows/win32_joystick.c")
	copyFile("thirdparty/glfw/src/win32_monitor.c", "dist/windows/win32_monitor.c")
	copyFile("thirdparty/glfw/src/win32_window.c", "dist/windows/win32_window.c")
	copyFile("thirdparty/glfw/src/wgl_context.c", "dist/windows/wgl_context.c")

	// Linux (X11)
	for _, file := range commonFiles {
		copyFile(file, "dist/linux/"+filepath.Base(file))
	}
	copyFile("thirdparty/glfw/src/x11_platform.h", "dist/linux/x11_platform.h")
	copyFile("thirdparty/glfw/src/xkb_unicode.h", "dist/linux/xkb_unicode.h")
	copyFile("thirdparty/glfw/src/x11_init.c", "dist/linux/x11_init.c")
	copyFile("thirdparty/glfw/src/x11_monitor.c", "dist/linux/x11_monitor.c")
	copyFile("thirdparty/glfw/src/x11_window.c", "dist/linux/x11_window.c")
	copyFile("thirdparty/glfw/src/xkb_unicode.c", "dist/linux/xkb_unicode.c")
	copyFile("thirdparty/glfw/src/glx_context.c", "dist/linux/glx_context.c")
	copyFile("thirdparty/glfw/src/linux_joystick.h", "dist/linux/linux_joystick.h")
	copyFile("thirdparty/glfw/src/linux_joystick.c", "dist/linux/linux_joystick.c")
	copyFile("thirdparty/glfw/src/posix_poll.h", "dist/linux/posix_poll.h")
	copyFile("thirdparty/glfw/src/posix_poll.c", "dist/linux/posix_poll.c")
	copyFile("thirdparty/glfw/src/posix_time.h", "dist/linux/posix_time.h")
	copyFile("thirdparty/glfw/src/posix_thread.h", "dist/linux/posix_thread.h")
	copyFile("thirdparty/glfw/src/posix_module.c", "dist/linux/posix_module.c")
	copyFile("thirdparty/glfw/src/posix_time.c", "dist/linux/posix_time.c")
	copyFile("thirdparty/glfw/src/posix_thread.c", "dist/linux/posix_thread.c")

	copyFile("thirdparty/glfw/include/GLFW/glfw3.h", "dist/include/GLFW/glfw3.h")
	copyFile("thirdparty/glfw/include/GLFW/glfw3native.h", "dist/include/GLFW/glfw3native.h")

	generateGlfwConstants()
}
