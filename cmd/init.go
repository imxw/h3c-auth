package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/imxw/h3c-auth/internal/pkg/fileutil"
)

const cfgPath = ".auth.yml"

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a config file and an automation script",
	Long:  "Initialize (h3cauth init) will create a default config file and an automation script",
	Run: func(cmd *cobra.Command, args []string) {
		err := createCfgInCwd(cmd, cfgPath)
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

func createCfgInCwd(cmd *cobra.Command, file string) error {

	if !fileutil.Exists(file) {
		f, err := os.Create(cfgPath)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.Write(genCfg(cmd, defaultCfg))
		return err
	} else {
		isForce, err := cmd.PersistentFlags().GetBool("force")
		cobra.CheckErr(err)
		if isForce {
			f, err := os.Create(cfgPath)
			if err != nil {
				return err
			}
			defer f.Close()
			_, err = f.Write(genCfg(cmd, defaultCfg))
			return err
		}
	}

	return nil
}

func genCfg(cmd *cobra.Command, template string) []byte {

	var err error
	var name, pwd string

	name, err = cmd.PersistentFlags().GetString("username")
	cobra.CheckErr(err)
	pwd, err = cmd.PersistentFlags().GetString("password")
	cobra.CheckErr(err)

	return []byte(fmt.Sprintf(defaultCfg, name, pwd))
}

func init() {
	userV := "username"
	pwdV := "password"
	initCmd.PersistentFlags().StringP(userV, "u", "USERNAME", "Your H3C account")
	initCmd.PersistentFlags().StringP(pwdV, "p", "PASSWORD", "Your H3C password")
	initCmd.PersistentFlags().BoolP("force", "f", false, "Force to create config file, will override your previous config (default false)")
}
