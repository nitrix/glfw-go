package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
)

func copyFile(src, dst string) error {
	_ = os.MkdirAll(filepath.Dir(dst), 0750)
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = srcFile.WriteTo(dstFile)
	if err != nil {
		return err
	}

	fmt.Printf("Copied file %s to %s\n", src, dst)

	return nil
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

func generateHeaders() {
	copyFile("thirdparty/glfw/include/GLFW/glfw3.h", "dist/include/GLFW/glfw3.h")
	copyFile("thirdparty/glfw/include/GLFW/glfw3native.h", "dist/include/GLFW/glfw3native.h")
	copyFile("thirdparty/glad/include/glad.h", "dist/include/glad/glad.h")
}

func generateGlfwConstants() {
	preamble := "package glfw\n"
	preamble += "\n"
	preamble += "// #cgo CFLAGS: -I../dist/include\n"
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
	generateHeaders()
	generateGlfwConstants()
}
