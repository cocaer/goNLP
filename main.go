package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/cocaer/goNLP/seg/hmm"
)

func main() {
	//		bhmm.HmmSaveTrainingFile()
	k := time.Now()
	m := hmm.NewModel()
	in, _ := os.Open("/home/begosu/gocode/src/github.com/cocaer/goNLP/test/pku_test.utf8")
	defer in.Close()
	out, err := os.Create("./test/test.utf8")
	if err != nil {
		fmt.Print("create file error")
	}
	defer out.Close()
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		s := scanner.Text()

		words := m.Cut(s)
		for _, word := range words {
			out.WriteString(word + " ")
		}
		out.WriteString("\n")
	}
	fmt.Println(time.Now().Sub(k))
}
