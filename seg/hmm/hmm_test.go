package hmm

import (
	"bufio"
	"os"
	"testing"
)

func TestHmm(t *testing.T) {
	m := NewModel()
	in, _ := os.Open("/b/homeegosu/gocode/src/github.com/cocaer/goNLP/data/msr_training.utf8")

	defer in.Close()
	out, _ := os.Create("../data/msr_out.txt")
	defer out.Close()
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		s := scanner.Text()
		words := m.Cut(s)
		for _, word := range words {
			out.WriteString(word)
		}
		out.WriteString("\n")
	}

}
