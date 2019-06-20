/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"wireguard-configer/randString"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var (
	user *string
	ip   *string
)
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a user.",
	Long: `add a user to your wireguard
 For example:
	wgt add -u me -s 10.0.0.4
`,
	Run: func(cmd *cobra.Command, args []string) {
		if *user == "" {
			*user = randString.RandStringBytes(8)
			fmt.Printf("%c[1;40;32m%s%c[0m", 0x1B, "[+] ", 0x1B)
			fmt.Println("System auto gen new username: ", *user)
		}
		if *ip == "" {
			fmt.Printf("%c[1;40;31m%s%c[0m", 0x1B, "[-] ", 0x1B)
			fmt.Errorf("Please enter ip info")
		}
		if *ip != "" && *user != "" {
			fmt.Printf("%c[1;40;31m%s%c[0m", 0x1B, "[-] ", 0x1B)
			fmt.Println("User: " + *user + "\nIP: " + *ip)
			// au := AddUser{*ip, }
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	user = addCmd.Flags().StringP("user", "u", "", "added wireguard username")
	ip = addCmd.Flags().StringP("src", "s", "", "added wireguard user ip")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
