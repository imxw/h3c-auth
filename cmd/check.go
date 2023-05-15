package cmd

import (
	"errors"
	"fmt"

	"github.com/imxw/h3c-auth/internal/pkg/netutil"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if the network is available",
	Long:  `Check if the network is available`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ret := netutil.IsNetOk()
		if ret {
			ipAddr, err := netutil.GetLocalIP()
			if err != nil {
				return err
			}
			fmt.Println("Network is available, and your local ip is: ", ipAddr)
			return nil
		}
		return errors.New("your network is not available")
	},
}
