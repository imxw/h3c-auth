package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/imxw/h3c-auth/internal/pkg/h3cauth"
	"github.com/imxw/h3c-auth/internal/pkg/netutil"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Auth the network connection of h3c",
	Long:  `Auth the network connection of h3c`,
	RunE: func(cmd *cobra.Command, args []string) error {
		localIp, err := netutil.GetLocalIP()
		if err != nil {
			return err
		}

		isIn := false

		nets := []string{"10.0.156.0/22", "10.0.44.0/22"}

		for _, v := range nets {
			if !netutil.IsIpInNet(localIp, v) {
				continue
			}
			isIn = true
			break
		}

		if !isIn {
			fmt.Println("你没有连公司的网络，无需认证")
			return nil
		} else {

			if netutil.IsNetOk() {
				fmt.Println("网络正常，无需认证")
				return nil
			}
			fmt.Println("Start to auth ...")
			err := h3cauth.Auth()
			if err != nil {
				fmt.Println(err)
				return nil
			}
			fmt.Println("Success!")
		}
		return nil
	},
}
