package cmd

import (
	"fmt"
	"strconv"

	"github.com/DarthXibalba/Rocinante/internal/fileio"
	"github.com/DarthXibalba/Rocinante/internal/imgproc"
	"github.com/spf13/cobra"
)

var scaleCmd = &cobra.Command{
	Use:   "scale [imgPath] [newWidth] [newHeigth]",
	Short: "Scale an image to new resolution",
	Long:  "Scale an image using the highest quality (non-performance method)",
	Args:  cobra.ExactArgs(3),
	RunE: func(_ *cobra.Command, args []string) error {
		// Load
		imgPath := args[0]
		img, err := imgproc.LoadImage(imgPath)
		if err != nil {
			return err
		}

		// Scale
		var newHeight, newWidth int
		newWidth, err = strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("error converting newWidth to integer: %s", err)
		}
		newHeight, err = strconv.Atoi(args[2])
		if err != nil {
			return fmt.Errorf("error converting newHeight to integer: %s", err)
		}
		scaledImg, err := imgproc.ScaleImage(img, newWidth, newHeight)
		if err != nil {
			return fmt.Errorf("error scaling image: %s", err)
		}

		// Save
		newImgPath := fileio.AddSuffixToFilename(imgPath, fmt.Sprintf("_scaled_%dx%d", newWidth, newHeight))
		err = imgproc.SaveImage(scaledImg, newImgPath)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(scaleCmd)
}
