package command

import "os/exec"

func Runner(flag string, args ...string) (string, error) {

    cmd := exec.Command(flag, args...)
    stdout, err := cmd.Output()

    if err != nil {
        return string(stdout), err
    }

    return string(stdout), nil
}