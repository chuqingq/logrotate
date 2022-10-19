package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("ERROR: please specify config file")
		os.Exit(-1)
	}

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("ERROR: read config file error: %v\n", err)
		os.Exit(-1)
	}

	l := &lumberjack.Logger{}
	err = json.Unmarshal(b, l)
	if err != nil {
		fmt.Printf("ERROR: parse config file error: %v\n", err)
		os.Exit(-1)
	}

	io.Copy(l, os.Stdin)
}
