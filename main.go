package main

import (
	"github.com/mengboy/img/cmd"
	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use:   "img",
	Short: "image compress",
}

func init() {
	rootCMD.AddCommand(cmd.ConvertCMD)
	rootCMD.AddCommand(cmd.QualityCMD)
	rootCMD.AddCommand(cmd.CompressCMD)
	rootCMD.AddCommand(cmd.VersionCMD)
}

func main() {
	_ = rootCMD.Execute()
}
