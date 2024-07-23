package config

import (
	"flag"
)

type Config struct {
	Lang string 
	Url  string 
}

type Options struct {
	Lang string
}

func options() Options {
	opt := Options{}
	flag.StringVar(&opt.Lang, "lang", "fa", "language for translate")
	flag.Parse()
	return opt
}

func getConfig() Config {
	opt := options()
	config := Config{
		Url: "http://translate.googleapis.com/translate_a/single?client=gtx&sl=auto&tl=%s&dt=t&q=%s",
		Lang: opt.Lang,
	}
	return config
}
