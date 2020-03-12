package main

import (
	"flag"
	"fmt"
	"github.com/keizo042/wiser-go"
	"os"
)

const (
	exitSuccess = 0
	exitFailure = -1
)

func main() {
	os.Exit(realMain(os.Args))
}

func realMain(argv []string) int {
	flag.Parse()
	env, err := wiser.NewEnv()
	if err != nil {
		fmt.Println(err)
		return exitFailure
	}
	return exitFailure
}
