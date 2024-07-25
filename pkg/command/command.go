package command

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

const pidLogFile string = ".translatorPID.log"

func Runner(flag string, args ...string) (string, error) {
	cmd := exec.Command(flag, args...)
	stdout, err := cmd.Output()
	if err != nil {
		return string(stdout), err
	}
	return string(stdout), nil
}

func CmdStart(flag string, args ...string) (int, error) {
	var pid int
	cmd := exec.Command(flag, args...)
	err := cmd.Start()
	if err != nil {
		return pid, err
	}
	pid = cmd.Process.Pid
	return pid, nil
}

func SavePidInFile(pid int) error {
	userHomePath, err := userHomeDir()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(userHomePath+"/"+pidLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%d\n", pid))
	if err != nil {
		return err
	}
	return nil
}

func GetPIDOnFile() ([]string, error) {
	var pidSlice []string
	userHomePath, err := userHomeDir()
	if err != nil {
		return pidSlice, err
	}
	file, err := os.Open(userHomePath + "/" + pidLogFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "\n" {
			continue
		}
		pidSlice = append(pidSlice, line)
	}
	return pidSlice, nil
}

func userHomeDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return homeDir, err
	}
	return homeDir, nil
}

func KillPid(pids []string) {
	for _, pid := range pids {
		Runner("kill", "-9", pid)
	}
}

func ClearLogFile() error {
	userHomePath, err := userHomeDir()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(userHomePath+"/"+pidLogFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	err = file.Truncate(0)
	if err != nil {
		return err
	}
	return nil
}
