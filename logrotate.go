package main 

import (
	// "log"
	// "time"
	"io"
	"io/ioutil"
	"os"
	"encoding/json"
	"fmt"
	// "lumberjack"
 "gopkg.in/natefinch/lumberjack.v2"
)

// To use lumberjack with the standard library's log package, just pass it into
// the SetOutput function when your application starts.
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
	/*
	l := &lumberjack.Logger{
		Filename:   "/mnt/d/temp/lumberjack/chuqq_test/myapp.log",
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     1,   // days
		Compress:   true, // disabled by default
	}
	log.SetOutput(l)
	*/
	io.Copy(l, os.Stdin)
}
