package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	VersionCMD = &cobra.Command{
		Use:   "version",
		Short: "version",
		Run:   version,
	}
)

func version(cmd *cobra.Command, args []string) {
	fmt.Println("1.0.0")
}
