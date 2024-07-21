package main

import (
	"fmt"

	"github.com/nopeEOF/translator/pkg/client"
	"github.com/nopeEOF/translator/pkg/command"
	"github.com/nopeEOF/translator/pkg/config"
	"github.com/nopeEOF/translator/pkg/translate"
)

func main() {
	config, err := config.NewConfig("config.json")
	if err != nil {
		panic(err)
	}
	client := client.NewClient(30)
	selectedClipboard, err := command.Runner("xsel", "-o")
	if err != nil {
		panic(err)
	}
	body, err := translate.Translate(selectedClipboard, config, client)
	if err != nil {
		panic(err)
	}

	body, _ = client.GetTranslateTextWithSplitBody(body)
	fmt.Println(body)
}
