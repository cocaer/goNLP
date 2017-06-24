package main

import (
	"fmt"

	"github.com/cocaer/goNLP/seg"
)

func main() {
	//seg.HmmSaveTraning()
	m := seg.NewHmmSeg() //基于hmm的优化实现
	fmt.Println(m.Cut("结婚的和尚未结婚的"))
}
