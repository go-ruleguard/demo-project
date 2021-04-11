package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

func imageColors(m dsl.Matcher) {
	m.Match(`color.Gray16{}`, `color.Gray16{0}`, `color.Gray16{$_: 0}`).
		Suggest(`color.Black`)

	m.Match(`color.Gray16{$x}`, `color.Gray16{$_: $x}`).
		Where(m["x"].Value.Int() == 0xffff).
		Suggest(`color.White`)

	m.Match(`color.Alpha16{}`, `color.Alpha16{0}`, `color.Alpha16{$_: 0}`).
		Suggest(`color.Transparent`)

	m.Match(`color.Alpha16{$x}`, `color.Alpha16{$_: $x}`).
		Where(m["x"].Value.Int() == 0xffff).
		Suggest(`color.Opaque`)
}

func imageZP(m dsl.Matcher) {
	m.Match(`image.ZP`).
		Report(`image.ZP is deprecated, use image.Point{} instead`).
		Suggest(`image.Point{}`)
}

func imagePt(m dsl.Matcher) {
	m.Match(`image.Pt(0, 0)`, `image.Point{0, 0}`).
		Report(`zero point should be written as image.Point{}`).
		Suggest(`image.Point{}`)

	m.Match(`image.Point{$x, $y}`).
		Report(`could use image.Pt() helper function`).
		Suggest(`image.Pt($x, $y`)
}

func imageRect(m dsl.Matcher) {
	m.Match(`image.Rectangle{Max: $max, Min: $min}`).
		Report(`Min field is mentioned before Max field`).
		Suggest(`image.Rectangle{Min: $min, Max: $max}`)

	m.Match(
		`image.Rectangle{image.Pt($x0, $y0), image.Pt($x1, $y1)}`,
		`image.Rectangle{Min: image.Pt($x0, $y0), Max: image.Pt($x1, $y1)}`,
	).Report(`could use image.Rect() helper function`).
		Suggest(`image.Rect($x0, $y0, $x1, $y1`)
}
