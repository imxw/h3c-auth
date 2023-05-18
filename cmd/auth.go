// Copyright 2023 Roy Xu <imxw1991@126.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/imxw/h3c-auth.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/go-playground/validator/v10"

	"github.com/imxw/h3c-auth/internal/pkg/h3cauth"
	"github.com/imxw/h3c-auth/internal/pkg/netutil"
	"github.com/imxw/h3c-auth/internal/pkg/notify"
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

		if err := checkCfg(); err != nil {
			return err
		}

		if viper.IsSet("netSegment") {

			nets := viper.GetStringSlice("netSegment")
			for _, v := range nets {
				if !netutil.IsIpInNet(localIp, v) {
					continue
				}
				isIn = true
				break
			}
		}

		if !isIn {
			notifyMsg("你没有连公司的网络，无需认证")
			return nil
		} else {

			if netutil.IsNetOk() {
				notifyMsg("网络正常，无需认证")
				return nil
			}
			cfg := h3cauth.Config{
				Username: viper.GetString("username"),
				Password: viper.GetString("password"),
				IpAddr:   viper.GetString("ipAddr"),
				Port:     viper.GetString("port"),
			}
			validate := validator.New()
			if err := validate.Struct(cfg); err != nil {
				return err
			}

			notifyMsg("Start to auth...")
			err := h3cauth.Auth(cfg)
			if err != nil {
				return err
			}
			notifyMsg("Success")
		}
		return nil
	},
}

func init() {

	userV := "username"
	pwdV := "password"
	authCmd.PersistentFlags().StringP(userV, "u", "", "Your H3C account")
	authCmd.PersistentFlags().StringP(pwdV, "p", "", "Your H3C password")
	err := viper.BindPFlag(userV, authCmd.PersistentFlags().Lookup(userV))
	cobra.CheckErr(err)
	err = viper.BindPFlag(pwdV, authCmd.PersistentFlags().Lookup(pwdV))
	cobra.CheckErr(err)
}

func notifyMsg(msg string) {
	if viper.GetBool("isNotify") {
		note := notify.NewNotification(msg)
		err := note.Push()
		cobra.CheckErr(err)
	}
	fmt.Println(msg)
}

func checkCfg() error {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return fmt.Errorf("%s is not exist in current directory, please execute \"h3cli init\" to generate one first", cfgPath)
		} else {
			return err
		}
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		return nil
	}
}
