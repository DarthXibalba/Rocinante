package cmd

import (
	"fmt"

	"github.com/DarthXibalba/Rocinante/internal/fileio"
	"github.com/DarthXibalba/Rocinante/internal/imgproc"
	"github.com/spf13/cobra"
)

var sharpenCmd = &cobra.Command{
	Use:   "sharpen [imgPath]",
	Short: "Sharpen an image",
	Long:  "Sharpen an image to improve enhance image features.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Load
		imgPath := args[0]
		img, err := imgproc.LoadImage(imgPath)
		if err != nil {
			return err
		}

		// Sharpen
		sharpenedImg, err := imgproc.SharpenImage(img)
		if err != nil {
			return fmt.Errorf("error sharpening image: %s", err)
		}

		// Save
		newImgPath := fileio.AddSuffixToFilename(imgPath, "_sharpened")
		err = imgproc.SaveImage(sharpenedImg, newImgPath)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(sharpenCmd)
}
