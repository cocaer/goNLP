package data

//
//给定一个已经分好词的文件，
//将它转换成BEMS序列
//
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CreateBEMSFile(input string, output string) {

	inputFile, err := os.Open(input)
	defer inputFile.Close()
	if err != nil {
		fmt.Println("Open FIle ", input, " Failed.")
		os.Exit(1)
	}

	outputFile, err := os.Create(output)
	defer outputFile.Close()
	if err != nil {
		fmt.Println("Create FIle ", output, " Failed.")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		for _, word := range words {
			sword := []rune(word)
			if len(sword) == 1 {
				outputFile.WriteString(string(sword) + " S\n")
			} else {
				outputFile.WriteString(string(sword[0]) + " B\n")

				for i := 1; i < len(sword)-1; i++ {
					outputFile.WriteString(string(sword[i]) + " M\n")
				}
				outputFile.WriteString(string(sword[len(sword)-1]) + " E\n")
			}
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Read error in ", input)
	}

}
