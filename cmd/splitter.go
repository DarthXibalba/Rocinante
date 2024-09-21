package cmd

import (
	"fmt"
	"image"
	"image/draw"

	"github.com/DarthXibalba/Rocinante/internal/fileio"
	"github.com/DarthXibalba/Rocinante/internal/imgproc"
	"github.com/spf13/cobra"
)

var (
	flag_splitterImgWidth  int
	flag_splitterImgHeight int
)

// splitAppendCmd represents the splitter command
var splitAppendCmd = &cobra.Command{
	Use:   "split [imgPath]",
	Short: "Split an image into parts",
	Long:  `Split an image into multiple parts based on specified criteria.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		imgPath := args[0]
		img, err := imgproc.LoadImage(imgPath)
		if err != nil {
			return err
		}
		cmd.Println("Loaded image:", imgPath)

		// Create a mirror image for the backside
		mirroredImg := imgproc.MirrorImage(img)

		// Calculate the bounds and the middle of the image
		origBounds := img.Bounds()
		origHeight := origBounds.Dy()
		origWidth := origBounds.Dx()
		origMiddle := origWidth / 2

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
		draw.Draw(newImg, image.Rect(0, 0, origMiddle, newHeight), img, image.Point{X: origMiddle, Y: 0}, draw.Src)
		draw.Draw(newImg, image.Rect(origMiddle, 0, origWidth+origMiddle, newHeight), mirroredImg, image.Point{}, draw.Src)
		draw.Draw(newImg, image.Rect(origWidth+origMiddle, 0, newWidth, newHeight), img, image.Point{}, draw.Src)

		cmd.Println("Split image")

		// Save images
		newImgPath := fileio.AddSuffixToFilename(imgPath, fmt.Sprintf("_cropped_%dx%d", newWidth, newHeight))
		newImgPath = fileio.ChangeFilenameExtension(newImgPath, ".png")
		err = imgproc.SaveImage(newImg, newImgPath)
		if err != nil {
			return err
		}

		cmd.Println("Saved image to:", newImgPath)
		return nil
	},
}

func init() {
	splitAppendCmd.Flags().IntVarP(&flag_splitterImgHeight, "height", "H", -1, "Sets the height (in pixels) of the resultant image. Default value will result in the output image height being the input image height.")
	splitAppendCmd.Flags().IntVarP(&flag_splitterImgWidth, "width", "W", -1, "Sets the width (in pixels) of the resultant image. Default value will result in the output image width being double of the input image width.")
	rootCmd.AddCommand(splitAppendCmd)
}
