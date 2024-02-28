package cmd

import (
	"fmt"
	"image"
	"image/draw"
	"path/filepath"
	"strings"

	"github.com/DarthXibalba/Image-Splitter/internal/imgproc"
	"github.com/spf13/cobra"
)

// splitAppendCmd represents the splitter command
var origSplitAppendCmd = &cobra.Command{
	Use:   "origSplit",
	Short: "Split an image into parts",
	Long:  `Split an image into multiple parts based on specified criteria.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		imgPath := args[0]
		img, err := imgproc.LoadImage(imgPath)
		if err != nil {
			return err
		}

		// Create a mirror image for the backside
		mirroredImg := imgproc.MirrorImage(img)

		// Calculate the bounds and the middle of the image
		bounds := img.Bounds()
		height := bounds.Dy()
		width := bounds.Dx()
		middle := width / 2

		// Create a new canvas with double the width to accommodate the RLRL pattern
		newWidth := width * 2
		newRect := image.Rect(0, 0, newWidth, height)
		newImg := image.NewRGBA(newRect)

		// Draw the 2nd image in RM(LR)L pattern
		draw.Draw(newImg, image.Rect(0, 0, middle, height), img, image.Point{X: middle, Y: 0}, draw.Src)
		draw.Draw(newImg, image.Rect(middle, 0, width+middle, height), mirroredImg, image.Point{}, draw.Src)
		draw.Draw(newImg, image.Rect(width+middle, 0, newWidth, height), img, image.Point{}, draw.Src)

		// Save images
		ext := filepath.Ext(imgPath)
		base := strings.TrimSuffix(imgPath, ext)
		newExt := ".png"
		suffix := fmt.Sprintf("_split_%dx%d", newWidth, height)

		newImgPath := base + suffix + newExt
		err = imgproc.SaveImage(newImg, newImgPath)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	// Here you can define flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// splitterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(origSplitAppendCmd)
}
