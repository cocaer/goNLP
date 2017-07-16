package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/cocaer/goNLP/seg"
)

func main() {
	in, _ := os.Open("/home/begosu/Documents/gocode/src/github.com/cocaer/goNLP/test/msr_test.utf8")
	defer in.Close()
	out, _ := os.Create("/home/begosu/Documents/gocode/src/github.com/cocaer/goNLP/test/out.utf8")
	defer out.Close()
	scaner := bufio.NewScanner(in)
	now := time.Now()
	seg.HmmSaveTraning()
	m := seg.NewHmmSeg()
	for scaner.Scan() {
		r := m.Cut(scaner.Text())
		for _, s := range r {
			out.WriteString(s)
			out.WriteString(" ")
		}
		out.WriteString("\n")
	}
	fmt.Println(time.Now().Sub(now))

}
