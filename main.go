package main

import (
	"flag"
	"fmt"
	"os"
	"plugin"
)

func main() {
	var file string
	var msg string
	flag.StringVar(&file, "f", "english.so", "plugin file")
	flag.StringVar(&msg, "m", "", "message")
	flag.Parse()

	p, err := plugin.Open(file)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(-1)
	}

	say, err := p.Lookup("SayHello")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(-1)
	}

	say.(func(string))(msg)
}
