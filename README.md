## Demo project that uses ruleguard

## Exploring the rules

Rules are located in [rules](/rules) package.

Run tests:

```bash
$ go test -v ./rules
=== RUN   TestRules
--- PASS: TestRules (1.00s)
PASS
```

Notes:

* Rules need `dsl` package, so it is present in the `go.mod`
* Since we have tests for the rules, we also need `golang.org/x/tools` package for `analysistest`

If you don't want to have `golang.org/x/tools` dependency in your project, then you should move
rules to a separate module and use it as a rules bundle in your main module.

## Running the rules

Run rules over the demo project:

```bash
$ ruleguard -c 0 -rules rules/rules.go ./mandelbrot
mandelbrot/main.go:36:12: imageRect: could use image.Rect() helper function (rules.go:44)
36		bounds := image.Rectangle{
mandelbrot/main.go:37:8: imagePt: zero point should be written as image.Point{} (rules.go:30)
37			Min: image.Pt(0, 0),
mandelbrot/main.go:41:11: imageColors: suggestion: color.Black (rules.go:8)
41		black := color.Gray16{0}
mandelbrot/main.go:42:48: imageZP: image.ZP is deprecated, use image.Point{} instead (rules.go:24)
```

Run rules with `golangci-lint`:

```bash
$ golangci-lint run ./mandelbrot
mandelbrot/main.go:36:12: ruleguard: could use image.Rect() helper function (gocritic)
	bounds := image.Rectangle{
	          ^
mandelbrot/main.go:37:8: ruleguard: zero point should be written as image.Point{} (gocritic)
		Min: image.Pt(0, 0),
		     ^
mandelbrot/main.go:41:11: ruleguard: suggestion: color.Black (gocritic)
	black := color.Gray16{0}
	         ^
mandelbrot/main.go:42:48: ruleguard: image.ZP is deprecated, use image.Point{} instead (gocritic)
	draw.Draw(b, bounds, image.NewUniform(black), image.ZP, draw.Src)
	                                              ^
```

See [.golangci.yml](.golangci.yml) config to see how to enable `ruleguard` for your golangci-lint.

Run rules with `gocritic`:

```bash
$ gocritic check -enable ruleguard -@ruleguard.rules rules/rules.go ./mandelbrot
./mandelbrot/main.go:36:12: ruleguard: could use image.Rect() helper function
./mandelbrot/main.go:42:48: ruleguard: image.ZP is deprecated, use image.Point{} instead
./mandelbrot/main.go:41:11: ruleguard: suggestion: color.Black
./mandelbrot/main.go:37:8: ruleguard: zero point should be written as image.Point{}
```

## Running the mandelbrot

```bash
go run ./mandelbrot

# or `go run ./mandelbrot/main.go
```

Enjoy the `mandelbrot.png`.
