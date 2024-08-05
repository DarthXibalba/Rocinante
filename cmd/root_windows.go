//go:build windows

package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   ".\\Rocinante.exe",
	Short: "Rocinante is a CLI tool for use image manipulation.",
	Long:  `Rocinante is a lightweight CLI tool use for a variety of image processing and image manipulation purposes.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Here you will define your flags and configuration settings.
}
