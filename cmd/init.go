// Copyright 2023 Roy Xu <imxw1991@126.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/imxw/h3c-auth.

package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"

	"github.com/imxw/h3c-auth/internal/pkg/fileutil"
)

const cfgPath = ".auth.yml"

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a config file and an automation script",
	Long:  "Initialize (h3cauth init) will create a default config file and an automation script",
	Run: func(cmd *cobra.Command, args []string) {
		err := genFilesInCwd(cmd)
		cobra.CheckErr(err)

	},
}

var defaultCfg = `username: %s # H3C Portal用户名
password: %s # H3C Portal密码
ipAddr: 10.0.100.20 # H3C Portal IP
port: ":8080" # H3C Portal Port
isNotify: true # 是否开启系统通知
netSegment:
  - "10.0.156.0/22"
  - "10.0.44.0/22"
`

const (
	startBat = `if not "%1"=="wkdxz" mshta vbscript:createobject("wscript.shell").run("""%~f0"" wkdxz",vbhide)(window.close)&&exit
@echo off
start "" /B cmd /C %userprofile%\auth.bat > output.log 2>&1
exit 
`
	authBat = `@echo off

setlocal enabledelayedexpansion
cd %userprofile%
:loop
set "time=%TIME%"
echo [%date% !time!] Running check command...
h3cauth.exe check

if %ERRORLEVEL% == 0 (
  set "time=%TIME%"
  echo [%date% !time!] Check command completed successfully. Sleeping for 5 seconds...
  timeout /t 5 >nul
) else (
  set "time=%TIME%"
  echo [%date% !time!] Check command failed with error %ERRORLEVEL%. Running auth command...
  (h3cauth.exe auth 2>&1) > auth_error.log
  type auth_error.log
  set "time=%TIME%"
  echo [%date% !time!] Auth command completed. Sleeping for 5 seconds...
  timeout /t 5 >nul
)

goto loop`
	startSh = `nohup bash auth.sh > output.log 2>&1 &`
	authSh  = `#!/bin/bash

while true
do
	# 检查网络连接
	h3cauth check
	if [ "$?" -eq "0" ]
	then
		# 如果网络连接正常，等待5秒钟后再次检查网络连接
		sleep 5
	else
		# 如果网络连接不正常，尝试连接网络
		echo "$(date '+%Y-%m-%d %H:%M:%S') Connecting to network..."
		h3cauth auth
		if [ "$?" -eq "0" ]
		then
			# 连接成功，等待5秒钟后再次检查网络连接
			echo "$(date '+%Y-%m-%d %H:%M:%S') Connected to network"
			sleep 5
		else
			# 连接失败，等待5秒钟后重试
			echo "$(date '+%Y-%m-%d %H:%M:%S') Failed to connect to network, retrying in 5 seconds..."
			sleep 5
		fi
	fi
done`
)

func genFilesInCwd(cmd *cobra.Command) error {

	isForce, err := cmd.PersistentFlags().GetBool("force")
	if err != nil {
		return err
	}
	username, err := cmd.PersistentFlags().GetString("username")
	if err != nil {
		return err
	}
	password, err := cmd.PersistentFlags().GetString("password")
	if err != nil {
		return err
	}

	templates := map[string]string{".auth.yml": defaultCfg, "start.bat": startBat, "auth.bat": authBat, "start.sh": startSh, "auth.sh": authSh}

	for k, v := range templates {
		if isForce || !fileutil.Exists(k) {
			if k == ".auth.yml" {
				v = fmt.Sprintf(v, username, password)
			}

			switch runtime.GOOS {
			case "windows":
				if strings.HasSuffix(k, ".sh") {
					continue
				}
			case "darwin":
				if strings.HasSuffix(k, ".bat") {
					continue
				}
			default:
				return fmt.Errorf("the current OS (%s) is not supported", runtime.GOOS)
			}

		}
		if err := os.WriteFile(k, []byte(v), 0666); err != nil {
			return err
		}
	}
	return nil

}

func init() {
	userV := "username"
	pwdV := "password"
	initCmd.PersistentFlags().StringP(userV, "u", "USERNAME", "Your H3C account")
	initCmd.PersistentFlags().StringP(pwdV, "p", "PASSWORD", "Your H3C password")
	initCmd.PersistentFlags().BoolP("force", "f", false, "Force to create config file, will override your previous config (default false)")
}
