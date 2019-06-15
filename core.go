package main

import (
	"fmt"
	"runtime"
)

var OSPackageTool = map[string]string{"darwin": "brew", "linux": "apt", "windows": ""}

func init() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

func main() {
	cmd := CmdInfo{`df -h`, "0"}
	if _, err := cmd.run(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cmd.String())
}
