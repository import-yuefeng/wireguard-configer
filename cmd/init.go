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

var (
	netInit    *string
	portInit   *string
	serverInit *string
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init wireguard",
	Long: `Init wireguard
For example:
	wgc init
	wgc init -i eth2
	wgc init -s 10.7.0.1
	wgc init -p 2345
`,
	Run: func(cmd *cobra.Command, args []string) {
		if exist, _ := pathExists(defaultConfPath); exist {
			fmt.Printf("%c[1;40;31m%s%c[0m", 0x1B, "[-] ", 0x1B)
			fmt.Println("Wireguard server config exist.")
			return
		}

		user := randString.RandStringBytes(8)
		init := initStruct{*netInit, *portInit, *serverInit}
		au := AddUser{"10.0.0.2", init.serverIP, init.listenPort, user}
		mkdirServerDir()
		genServerKeyFunc()
		au.genClientKeyFunc()
		init.genServerConfig()
		au.genClientConfig()
	},
}

type initStruct struct {
	netInterface string
	listenPort   string
	serverIP     string
}

func init() {

	rootCmd.AddCommand(initCmd)
	netInit = initCmd.Flags().StringP("interface", "i", "eth0", "net interface name")
	portInit = initCmd.Flags().StringP("port", "p", "4096", "wireguard listen port")
	serverInit = initCmd.Flags().StringP("src", "s", "10.0.0.1", "init wireguard server ip")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// func initServerConfig(serverIP *string, netInterface *string, listenPort *string) string {

// }
