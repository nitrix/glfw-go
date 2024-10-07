# glfw-go

A binding for glfw in Go.

## Example

![example.png](example.png)

Try the example with `go run github.com/nitrix/glfw-go/example@latest`.

The sources for it are located here [example/example.go](example/example.go).

## Progress

The bindings are written manually on an "as-needed" basis while the constants are scraped by a generator.
The GLFW API is kept intentionally compatible with `github.com/go-gl/glfw`. Report any deviation.

## Usage

Either import directly `github.com/nitrix/glfw-go/glfw` or if you have transitive dependencies, then you'll want to add `replace github.com/go-gl/glfw => github.com/nitrix/glfw-go/glfw` to your `go.mod` file.

## License

This is free and unencumbered software released into the public domain. See the [UNLICENSE](UNLICENSE) file for more details.
