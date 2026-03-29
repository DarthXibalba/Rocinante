package imgproc

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

// CenterImageOnCanvas draws img centered on a solid-color canvas.
// The input image must fit entirely within the canvas.
func CenterImageOnCanvas(img image.Image, canvasWidth, canvasHeight int, bg color.Color) (image.Image, error) {
	if img == nil {
		return nil, fmt.Errorf("input image cannot be nil")
	}
	if canvasWidth <= 0 || canvasHeight <= 0 {
		return nil, fmt.Errorf("canvas dimensions must be positive")
	}

	imgBounds := img.Bounds()
	imgWidth := imgBounds.Dx()
	imgHeight := imgBounds.Dy()

	if imgWidth <= 0 || imgHeight <= 0 {
		return nil, fmt.Errorf("input image dimensions must be positive")
	}
	if imgWidth > canvasWidth || imgHeight > canvasHeight {
		return nil, fmt.Errorf("input image (%dx%d) does not fit inside canvas (%dx%d)", imgWidth, imgHeight, canvasWidth, canvasHeight)
	}

	dst := image.NewRGBA(image.Rect(0, 0, canvasWidth, canvasHeight))
	draw.Draw(dst, dst.Bounds(), &image.Uniform{C: bg}, image.Point{}, draw.Src)

	offsetX := (canvasWidth - imgWidth) / 2
	offsetY := (canvasHeight - imgHeight) / 2
	targetRect := image.Rect(offsetX, offsetY, offsetX+imgWidth, offsetY+imgHeight)
	draw.Draw(dst, targetRect, img, imgBounds.Min, draw.Over)

	return dst, nil
}
