package main

import (
	"strings"

	"github.com/nopeEOF/translator/pkg/client"
	"github.com/nopeEOF/translator/pkg/command"
	"github.com/nopeEOF/translator/pkg/config"
	"github.com/nopeEOF/translator/pkg/translate"
)

func main() {
	pids, err := command.GetPIDOnFile()
    if err != nil {
        command.Runner("kdialog", "--msgbox", err.Error())
		return
    }
    command.KillPid(pids)
	err = command.ClearLogFile()
	if err != nil {
		command.Runner("kdialog", "--msgbox", err.Error())
		return
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
