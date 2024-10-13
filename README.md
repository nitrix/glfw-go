# glfw-go

A binding for glfw in Go.

## Example

![example.png](example.png)

Try the example with `go run github.com/nitrix/glfw-go/example@latest`.

The sources for it are located here [example/example.go](example/example.go).

## Usage

Either import directly `github.com/nitrix/glfw-go` or if you have transitive dependencies, then you'll want to add `replace github.com/go-gl/glfw => github.com/nitrix/glfw-go` to your `go.mod` file. For that reason, the API is kept intentionally compatible with `github.com/go-gl/glfw`. Report any deviations.

## License

This is free and unencumbered software released into the public domain. See the [UNLICENSE](UNLICENSE) file for more details.
