package target

import (
	"image"
	"image/color"
)

func testImageColor() {
	_ = color.Gray16{}     // want `\Qsuggestion: color.Black`
	_ = color.Gray16{0}    // want `\Qsuggestion: color.Black`
	_ = color.Gray16{Y: 0} // want `\Qsuggestion: color.Black`
	_ = color.Black        // OK

	_ = color.Alpha16{}     // want `\Qsuggestion: color.Transparent`
	_ = color.Alpha16{0}    // want `\Qsuggestion: color.Transparent`
	_ = color.Alpha16{A: 0} // want `\Qsuggestion: color.Transparent`
	_ = color.Transparent   // OK

	_ = color.Gray16{0xffff}    // want `\Qsuggestion: color.White`
	_ = color.Gray16{65535}     // want `\Qsuggestion: color.White`
	_ = color.Gray16{Y: 0xffff} // want `\Qsuggestion: color.White`
	_ = color.Gray16{Y: 65535}  // want `\Qsuggestion: color.White`
	_ = color.White             // OK

	_ = color.Alpha16{0xffff}    // want `\Qsuggestion: color.Opaque`
	_ = color.Alpha16{65535}     // want `\Qsuggestion: color.Opaque`
	_ = color.Alpha16{A: 0xffff} // want `\Qsuggestion: color.Opaque`
	_ = color.Alpha16{A: 65535}  // want `\Qsuggestion: color.Opaque`
	_ = color.Opaque             // OK
}

func testImageZP() {
	_ = image.ZP // want `\Qimage.ZP is deprecated, use image.Point{} instead`

	_ = image.Point{} // OK
}

func testImagePt() {
	_ = image.Point{0, 0} // want `\Qzero point should be written as image.Point{}`
	_ = image.Pt(0, 0)    // want `\Qzero point should be written as image.Point{}`

	_ = image.Point{1, 2}       // want `\Qcould use image.Pt() helper function`
	_ = image.Point{X: 1, Y: 2} // want `\Qcould use image.Pt() helper function`

	_ = image.Point{} // OK
}

func testImageRect() {
	_ = image.Rectangle{image.Pt(1, 1), image.Pt(2, 2)}           // want `\Qcould use image.Rect() helper function`
	_ = image.Rectangle{Min: image.Pt(1, 1), Max: image.Pt(2, 2)} // want `\Qcould use image.Rect() helper function`

	_ = image.Rectangle{Max: image.Pt(1, 1), Min: image.Pt(2, 2)} // want `\QMin field is mentioned before Max field`

	_ = image.Rectangle{Min: image.Pt(1, 1)} // OK: only 1 point is used
	_ = image.Rectangle{Max: image.Pt(2, 2)} // OK: only 1 point is used
	_ = image.Rectangle{}                    // OK: zero value

	_ = image.Rect(1, 1, 2, 2) // OK
}
