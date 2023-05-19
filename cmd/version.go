// Copyright 2023 Roy Xu <ixw1991@126.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/imxw/h3c-auth.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/imxw/h3c-auth/internal/pkg/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of h3cli",
	Long:  `All software has versions. This is h3cli's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Version)
	},
}
