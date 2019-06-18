package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"wireguard-configer/runCommand"
)

const (
	forwardIPV4     string = `echo "net.ipv4.ip_forward = 1" >> /etc/sysctl.conf`
	mkdirWireguard  string = `mkdir -p /etc/wireguard && chmod 0777 /etc/wireguard && cd /etc/wireguard`
	mkdirClientPath string = `/etc/wireguard/Client/`
	genServerKey    string = `wg genkey | tee /etc/wireguard/server_privatekey | wg pubkey > /etc/wireguard/server_publickey`
)

var (
	netInterface *string
	listenPort   *string
	serverIP     *string
)

var (
	defaultConfPath  = "/etc/wireguard/wg0.conf"
	initServerConfig = `
	echo "
[Interface]
	PrivateKey = $(cat /etc/wireguard/server_privatekey)
	Address = 
	PostUp = iptables -A FORWARD -i wg0 -j ACCEPT; iptables -A FORWARD -o wg0 -j ACCEPT; iptables -t nat -A POSTROUTING -o  
	PostDown = iptables -D FORWARD -i wg0 -j ACCEPT; iptables -D FORWARD -o wg0 -j ACCEPT; iptables -t nat -D POSTROUTING -o 
	ListenPort = 
	DNS = 8.8.8.8
	MTU = 1420 " > /etc/wireguard/wg0.conf`
	ListenPort    = "ListenPort = "
	PostUp        = `iptables -A FORWARD -i wg0 -j ACCEPT; iptables -A FORWARD -o wg0 -j ACCEPT; iptables -t nat -A POSTROUTING -o `
	PostStringEnd = ` -j MASQUERADE`
	PostDown      = `iptables -D FORWARD -i wg0 -j ACCEPT; iptables -D FORWARD -o wg0 -j ACCEPT; iptables -t nat -D POSTROUTING -o `

	genClientKey     = `wg genkey | tee /etc/wireguard/Client/client_privatekey | wg pubkey > /etc/wireguard/Client/client_publickey`
	initClientConfig = `echo "
	[Interface]
	  PrivateKey = $(cat client_privatekey)
	  Address = 10.0.0.2/24
	  DNS = 8.8.8.8
	  MTU = 1420
	
	[Peer]
	  PublicKey = $(cat server_publickey)
	  Endpoint = 1.2.3.4:50814
	  AllowedIPs = 0.0.0.0/0, ::0/0
	  PersistentKeepalive = 25 " > client.conf
	`
)

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func genClientKeyFunc(clientConf string) (bool, error) {
	genClientKey = strings.Replace(genClientKey, "client_privatekey", clientConf+"_privatekey", -1)
	genClientKey = strings.Replace(genClientKey, "client_publickey", clientConf+"_publickey", -1)

	runInit := runCommand.CmdInfo{genClientKey, "0", 0}
	if _, err := runInit.Exec(); err != nil {
		fmt.Printf("%c[1;40;31m%s%c[0m", 0x1B, "[-] ", 0x1B)
		fmt.Println(err)
		return false, errors.New("genClientKeyError")
	}
	return true, nil
}

func genServerKeyFunc() (bool, error) {
	runInit := runCommand.CmdInfo{genServerKey, "0", 0}
	if _, err := runInit.Exec(); err != nil {
		fmt.Printf("%c[1;40;31m%s%c[0m", 0x1B, "[-] ", 0x1B)
		fmt.Println(err)
		return false, errors.New("genServerKeyError")
	}
	return true, nil
}

func mkdirServerDir() {
	runInit := runCommand.CmdInfo{mkdirWireguard, "0", 0}
	if _, err := runInit.Exec(); err != nil {
		fmt.Printf("%c[1;40;31m%s%c[0m", 0x1B, "[-] ", 0x1B)
		fmt.Println(err)
		return
	}
	runInit = runCommand.CmdInfo{forwardIPV4, "0", 0}
	if _, err := runInit.Exec(); err != nil {
		fmt.Printf("%c[1;40;31m%s%c[0m", 0x1B, "[-] ", 0x1B)
		fmt.Println(err)
		return
	}
	os.Mkdir(mkdirClientPath, 0777)

}

func genClientConfig() {

}

func genServerConfig() {

	initServerConfig = strings.Replace(initServerConfig, "Address = ", "Address = "+*serverIP, -1)
	initServerConfig = strings.Replace(initServerConfig, PostUp, PostUp+*netInterface+PostStringEnd, -1)
	initServerConfig = strings.Replace(initServerConfig, PostDown, PostDown+*netInterface+PostStringEnd, -1)
	initServerConfig = strings.Replace(initServerConfig, ListenPort, ListenPort+*listenPort, -1)
	if *serverIP == "" || *listenPort == "" || *netInterface == "" {
		fmt.Printf("%c[1;40;31m%s%c[0m", 0x1B, "[-] ", 0x1B)
		fmt.Println("Missing required arguments")
		return
	}
	runInit := runCommand.CmdInfo{initServerConfig, "0", 0}
	if _, err := runInit.Exec(); err != nil {
		fmt.Printf("%c[1;40;31m%s%c[0m", 0x1B, "[-] ", 0x1B)
		fmt.Println(err)
		return
	}
}
