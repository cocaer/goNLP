package main

import (
	"fmt"

	"github.com/cocaer/goNLP/seg/bhmm"
)

func main() {
	//		bhmm.HmmSaveTrainingFile()
	// k := time.Now()
	// m := hmm.NewModel()
	// in, _ := os.Open("/home/begosu/gocode/src/github.com/cocaer/goNLP/test/pku_test.utf8")
	// defer in.Close()
	// out, err := os.Create("./test/test.utf8")
	// if err != nil {
	// 	fmt.Print("create file error")
	// }
	// defer out.Close()
	// scanner := bufio.NewScanner(in)
	// for scanner.Scan() {
	// 	s := scanner.Text()

	// 	words := m.Cut(s)
	// 	for _, word := range words {
	// 		out.WriteString(word + " ")
	// 	}
	// 	out.WriteString("\n")
	// }
	// fmt.Println(time.Now().Sub(k))
	m := bhmm.NewModel() //基于hmm的优化实现
	//m := hmm.NewModel() //纯hmm实现，没有优化
	fmt.Println(m.Cut("结婚的和尚未结婚的"))
}
