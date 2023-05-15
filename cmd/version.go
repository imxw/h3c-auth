package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/imxw/h3c-auth/internal/pkg/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of h3cauth",
	Long:  `All software has versions. This is h3cauth's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Version)
	},
}
