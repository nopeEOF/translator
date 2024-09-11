package config

import (
	"flag"
)

type Config struct {
	Lang      string
	Url       string
	Direction string
	FontSize  int
}

type Options struct {
	Language  string
	Direction string
	FontSize  int
}

func options() Options {
	opt := Options{}
	flag.StringVar(&opt.Language, "lang", "en", "language for translate")
	flag.StringVar(&opt.Direction, "dir", "ltr", "The dir attribute specifies the text direction of the element's content")
	flag.IntVar(&opt.FontSize, "font", 4, "The font size in points")
	flag.Parse()
	return opt
}

func getConfig() Config {
	opt := options()
	config := Config{
		Url:       "http://translate.googleapis.com/translate_a/single?client=gtx&sl=auto&tl=%s&dt=t&q=%s",
		Lang:      opt.Language,
		Direction: opt.Direction,
		FontSize:  opt.FontSize,
	}
	return config
}
