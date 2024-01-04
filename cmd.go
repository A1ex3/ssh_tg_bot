package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type SshInfo struct {
	RemoteIpPort string
	RemoteIp     string
	HostIp string
	UserName     string
}

func (sshInfo *SshInfo) Get() error {
	ipCmd := exec.Command("bash", "-c", "echo $SSH_CONNECTION | awk -F' ' '{print $1}'")
	ipHostCmd := exec.Command("bash", "-c", "echo $SSH_CONNECTION | awk -F' ' '{print $3}'")
	portCmd := exec.Command("bash", "-c", "echo $SSH_CONNECTION | awk -F' ' '{print $4}'")
	userNameCmd := exec.Command("whoami")

	stdoutIp, errIp := ipCmd.Output()
	stdoutIpHost, errIpHost := ipHostCmd.Output()
	stdoutPort, errPort := portCmd.Output()
	stdoutUserName, errUserName := userNameCmd.Output()

	if errIp != nil || errPort != nil || errUserName != nil || errIpHost != nil {
		return fmt.Errorf("errIp: %s\nerrPort: %s\nerrUserName: %s\nerrIpHost: %s", errIp, errPort, errUserName, errIpHost)
	}

	sshInfo.RemoteIp = strings.TrimSpace(string(stdoutIp))
	sshInfo.HostIp = strings.TrimSpace(string(stdoutIpHost))
	sshInfo.RemoteIpPort = strings.TrimSpace(string(stdoutPort))
	sshInfo.UserName = strings.TrimSpace(string(stdoutUserName))
	return nil
}
