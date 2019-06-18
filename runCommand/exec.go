package runCommand

import (
	"io/ioutil"
	"os/exec"
	"syscall"
)

type CmdInfo struct {
	Command  string
	Result   string
	ExitCode interface{}
}

func (cmdinfo *CmdInfo) String() string {
	return cmdinfo.Result
}

func (cmdinfo *CmdInfo) Exec() (string, error) {

	cmd := exec.Command("/bin/bash", "-c", cmdinfo.Command)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "Error: can not obtain stdout pipe\n", err
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		return "ExecError\n", err
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "ReadAllStdoutError", err
	}

	if err := cmd.Wait(); err != nil {
		return "WairError", err
	}
	cmdinfo.Result = string(bytes)
	cmdinfo.ExitCode = cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	return cmdinfo.Result, nil
}
