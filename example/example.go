package main

import (
	"runtime"

	"github.com/nitrix/glfw-go"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Visible, glfw.False)

	window, err := glfw.CreateWindow(1280, 720, "Example", nil, nil)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode glfw.Scancode, action glfw.Action, mods glfw.ModifierKey) {
		if key == glfw.KeyEscape && action == glfw.Press {
			w.SetShouldClose(true)
		}
	})

	window.Centerize()
	window.Show()

	for !window.ShouldClose() {
		glfw.PollEvents()
		window.SwapBuffers()
	}
}
