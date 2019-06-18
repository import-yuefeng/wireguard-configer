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

	"github.com/spf13/cobra"
)

var deluser *string

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "For wireguard del a user",
	Long: `For wireguard del a user
For example:
	wgc del -u me
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if *deluser == "" {
			fmt.Printf("%c[1;40;31m%s%c[0m", 0x1B, "[-] ", 0x1B)
			fmt.Errorf("Please enter username.")
			cmd.Usage()
		} else {
			fmt.Printf("%c[1;40;31m%s%c[0m", 0x1B, "[-] ", 0x1B)
			fmt.Println("User: " + *deluser + " deleted.")
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
	deluser = delCmd.Flags().StringP("user", "u", "", "del a user")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
