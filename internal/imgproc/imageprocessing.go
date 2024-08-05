package imgproc

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
	"path/filepath"

	"golang.org/x/image/draw"
	"golang.org/x/image/webp"
)

// Load the image from the provided file path
func LoadImage(filePath string) (image.Image, error) {
	fmt.Println("Loading image:", filePath)
	var img image.Image
	file, err := os.Open(filePath)
	if err != nil {
		return img, fmt.Errorf("failed to open the file: %v", err)
	}
	defer file.Close()

	if filepath.Ext(filePath) == ".webp" {
		img, err = webp.Decode(file)
	} else {
		img, _, err = image.Decode(file)
	}

	if err != nil {
		return img, fmt.Errorf("failed to decode the image: %v", err)
	}
	return img, nil
}

// saveImage saves the given image to a file in PNG format
func SaveImage(img image.Image, filePath string) error {
	fmt.Println("Saving image:", filePath)
	outFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Encode and save the image as a PNG
	fmt.Println("Saved image!")
	return png.Encode(outFile, img)
}

// ScaleImage scales the given image to the specified dimensions and returns the scaled image
func ScaleImage(img image.Image, x, y int) (image.Image, error) {
	// Check for non-positive dimensions
	if x <= 0 || y <= 0 {
		return nil, fmt.Errorf("dimensions must be positive")
	}

	// Create a new empty image of the specified size
	dst := image.NewRGBA(image.Rect(0, 0, x, y))

	// Scale the image using the CatmullRom scaler for the highest quality
	draw.CatmullRom.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

	fmt.Println("Image scaled with high quality to:", x, "x", y)
	return dst, nil
}

// Creates a horizontally mirrored version of the given image
func MirrorImage(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Create a new empty image with the same size
	mirroredImg := image.NewRGBA(image.Rect(0, 0, width, height))

	// Copy each pixel from the original image to the mirrored position
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			mirroredPixel := img.At(width-x-1, y)
			mirroredImg.Set(x, y, mirroredPixel)
		}
	}

	return mirroredImg
}
