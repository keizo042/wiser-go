package main

import (
	"flag"
	"fmt"
	"github.com/keizo042/wiser-go"
	"os"
)

var (
	CompressionMethod       = flag.String("c", "", "")
	WikipediaDumpXML        = flag.String("x", "", "")
	SearchQuery             = flag.String("q", "", "")
	maxIndexCount           = flag.Int("m", -1, "")
	IIBufferUpdateThreshold = flag.Int("t", 0, "")
	EnablePharseSearch      = flag.Bool("s", false, "")
)

const (
	exitSuccess = 0
	exitFailure = -1
)

func main() {
	flag.Parse()
	os.Exit(realMain(flag.Args()))
}

func realMain(argv []string) int {
	if len(argv) <= 0 {
		return exitFailure
	}
	if WikipediaDumpXML != nil {
		if _, err := os.Stat(*WikipediaDumpXML); err != nil {
			fmt.Println(err)
			return exitFailure
		}
	}

	dbPath := argv[1]
	env, err := wiser.NewEnv(
		*IIBufferUpdateThreshold,
		*EnablePharseSearch,
		dbPath,
	)
	if err != nil {
		fmt.Println(err)
		return exitFailure
	}
	defer env.Close()

	var method CompressionMethod
	if CompressionMethod != nil {
		method = NewCompressionMethod(*CompressionMethod)
	}
	if err := env.Index(method); err != nil {
		fmt.Println(err)
		return exitFailure
	}
	return exitFailure
}
