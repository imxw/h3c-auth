// Copyright 2023 Roy Xu <ixw1991@126.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/imxw/h3c-auth.

package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/imxw/h3c-auth/internal/pkg/netutil"
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
