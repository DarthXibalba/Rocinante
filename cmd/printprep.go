package cmd

import (
	"fmt"
	"image/color"

	"github.com/DarthXibalba/Rocinante/internal/fileio"
	"github.com/DarthXibalba/Rocinante/internal/imgproc"
	"github.com/spf13/cobra"
)

var (
	flagPrintPrepCanvasWidth  int
	flagPrintPrepCanvasHeight int
	flagPrintPrepImageSize    int
)

var printPrepCmd = &cobra.Command{
	Use:   "printprep [imgPath]",
	Short: "Center a square photo on a print canvas",
	Long:  "Scales a square image to the desired print size and centers it on a larger print canvas. Useful for placing a 3x3 image in the center of a 4x6 print.",
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		imgPath := args[0]

		img, err := imgproc.LoadImage(imgPath)
		if err != nil {
			return err
		}

		bounds := img.Bounds()
		if bounds.Dx() != bounds.Dy() {
			return fmt.Errorf("printprep expects a square input image, got %dx%d", bounds.Dx(), bounds.Dy())
		}

		scaledImg, err := imgproc.ScaleImage(img, flagPrintPrepImageSize, flagPrintPrepImageSize)
		if err != nil {
			return fmt.Errorf("error scaling image: %w", err)
		}

		outputImg, err := imgproc.CenterImageOnCanvas(
			scaledImg,
			flagPrintPrepCanvasWidth,
			flagPrintPrepCanvasHeight,
			color.White,
		)
		if err != nil {
			return fmt.Errorf("error centering image on canvas: %w", err)
		}

		newImgPath := fileio.AddSuffixToFilename(
			imgPath,
			fmt.Sprintf("_printready_%dx%d_img%d", flagPrintPrepCanvasWidth, flagPrintPrepCanvasHeight, flagPrintPrepImageSize),
		)
		err = imgproc.SaveImage(outputImg, newImgPath)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	printPrepCmd.Flags().IntVarP(&flagPrintPrepCanvasWidth, "canvas-width", "W", 1800, "Canvas width in pixels. 4x6 at 300 DPI is 1800.")
	printPrepCmd.Flags().IntVarP(&flagPrintPrepCanvasHeight, "canvas-height", "H", 1200, "Canvas height in pixels. 4x6 at 300 DPI is 1200.")
	printPrepCmd.Flags().IntVarP(&flagPrintPrepImageSize, "image-size", "S", 900, "Square image size in pixels. 3x3 at 300 DPI is 900.")
	rootCmd.AddCommand(printPrepCmd)
}
