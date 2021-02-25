package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"time"
)

//go:embed resources/ofuton.txt
var ofutonTxt []byte

var (
	stdOut = bufio.NewWriter(os.Stdout)
)

func clear() {
	// https://www.grapecity.com/developer/support/powernews/column/clang/047/page02.htm
	fmt.Fprint(stdOut, "\x1b[2J")
	fmt.Fprint(stdOut, "\x1b[1;1H")
	stdOut.Flush()
}

func main() {
	clear()
	for {
		scanner := bufio.NewScanner(bytes.NewBuffer(ofutonTxt))
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			time.Sleep(200 * time.Millisecond)
			clear()
		}
	}
}
