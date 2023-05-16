// Copyright 2023 Roy Xu <imxw1991@126.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/imxw/h3c-auth.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "h3cauth",
		Short: "A cmd for h3c auth",
		Long:  `h3cauth is a command-line tool for h3c auth.`,
		Example: `  First, initialize a config using "h3cauth init -u USERNAME -p PASSWORD".
  Then connect to your network using "h3cauth auth"`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(authCmd)
	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(initCmd)

	authCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./.auth.yml)")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".auth")
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("auth")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
