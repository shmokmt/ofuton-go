package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/gobuffalo/packr/v2"
)

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
	f, err := packr.NewBox("./resources").Open("ofuton.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	clear()
	for {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			time.Sleep(200 * time.Millisecond)
			clear()
		}
	}
}
