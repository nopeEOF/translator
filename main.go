package main

import (
	"strings"

	"github.com/nopeEOF/translator/pkg/client"
	"github.com/nopeEOF/translator/pkg/command"
	"github.com/nopeEOF/translator/pkg/config"
	"github.com/nopeEOF/translator/pkg/translate"
)

func killPidAndClearLogFile() error {
	// kill all kdialog window with process id saved in log file
	pids, err := command.GetPIDOnFile()
    if err != nil {
		return err
    }
    command.KillPid(pids)

	// clear all pid in log file
	err = command.ClearLogFile()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := killPidAndClearLogFile()
	if err != nil {
		command.Runner("kdialog", "--msgbox", err.Error())
	}

	config := config.NewConfig()
	client := client.NewClient(5)
	selectedClipboard, err := command.Runner("xsel", "-o")
	selectedClipboard = strings.Replace(selectedClipboard, "\n", " ", -1)

	if err != nil {
		command.Runner("kdialog", "--msgbox", err.Error())
		return
	}
	body, err := translate.Translate(selectedClipboard, config, client)
	if err != nil {
		command.Runner("kdialog", "--msgbox", "check internet connection")
		return
	}
	body, err = client.GetTranslateTextWithSplitBody(body)
	if err != nil {
		command.Runner("kdialog", "--msgbox", err.Error())
		return
	}
	body = client.KdialogMessageBody(body)
	pid, err := command.CmdStart("kdialog", "--msgbox", body)
	command.SavePidInFile(pid)
}
