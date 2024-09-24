package cmd

import (
	"fmt"
	"image"
	"image/draw"

	"github.com/DarthXibalba/Rocinante/internal/fileio"
	"github.com/DarthXibalba/Rocinante/internal/imgproc"
	"github.com/spf13/cobra"
)

// splitAppendCmd represents the splitter command
var duplicateAppendCmd = &cobra.Command{
	Use:   "duplicate [imgPath]",
	Short: "duplicate an image",
	Long:  `Duplicate an image into adjacent parts`,
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		imgPath := args[0]
		img, err := imgproc.LoadImage(imgPath)
		if err != nil {
			return err
		}

		// Calculate the bounds and the middle of the image
		origBounds := img.Bounds()
		origHeight := origBounds.Dy()
		origWidth := origBounds.Dx()

		newHeight := origHeight
		newWidth := origWidth * 2
		if flag_splitterImgHeight > 0 {
			newHeight = flag_splitterImgHeight
		}
		if flag_splitterImgWidth > 0 {
			newWidth = flag_splitterImgWidth
		}

		// Create a new canvas with double the width to accommodate the RLRL pattern
		newRect := image.Rect(0, 0, newWidth, newHeight)
		newImg := image.NewRGBA(newRect)

		// Draw the 2nd image in RM(LR)L pattern
		draw.Draw(newImg, image.Rect(0, 0, origWidth, origHeight), img, image.Point{}, draw.Src)
		draw.Draw(newImg, image.Rect(origWidth, 0, 2*origWidth, origHeight), img, image.Point{}, draw.Src)

		//draw.Draw(newImg, image.Rect(0, 0, origMiddle, newHeight), img, image.Point{X: origMiddle, Y: 0}, draw.Src)
		//draw.Draw(newImg, image.Rect(origMiddle, 0, origWidth+origMiddle, newHeight), mirroredImg, image.Point{}, draw.Src)
		//draw.Draw(newImg, image.Rect(origWidth+origMiddle, 0, newWidth, newHeight), img, image.Point{}, draw.Src)

		// Save images
		newImgPath := fileio.AddSuffixToFilename(imgPath, fmt.Sprintf("_cropped_%dx%d", newWidth, newHeight))
		newImgPath = fileio.ChangeFilenameExtension(newImgPath, ".png")
		err = imgproc.SaveImage(newImg, newImgPath)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(duplicateAppendCmd)
}
