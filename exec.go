package main

import (
	"io/ioutil"
	"os/exec"
)

type CmdInfo struct {
	command string
	result  string
}

func (cmdinfo *CmdInfo) String() string {
	return cmdinfo.command + "result is: \n" + cmdinfo.result
}

func (cmdinfo *CmdInfo) run() (string, error) {

	cmd := exec.Command("/bin/bash", "-c", cmdinfo.command)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "Error: can not obtain stdout pipe\n", err
	}

	if err := cmd.Start(); err != nil {
		return "EXECError\n", err
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "ReadAllStdoutError", err
	}

	if err := cmd.Wait(); err != nil {
		return "WairError", err
	}
	cmdinfo.result = string(bytes)
	return cmdinfo.result, nil
}
