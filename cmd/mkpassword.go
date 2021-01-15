package cmd

import (
	"ff/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var mkPwdCmd = &cobra.Command{
	Use:   "mkpassword",
	Short: "use this command to generate one random password",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		plainText, _ := utils.GenRandomStr(16)
		pwd := new(utils.Password)
		pwd.MkPassword(plainText)
		fmt.Println("明文:",pwd.PlainText)
		fmt.Println("密文:",pwd.CipherText)
	},
}
