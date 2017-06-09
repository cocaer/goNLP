package hmm

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	seg "github.com/cocaer/goNLP/seg/config"
)

//Feature 一个汉字一个Feature
type Feature struct {
	Count   int
	BEMS    [SUM_STATUS]int
	BEMSPro [SUM_STATUS]float64
}

var ma = map[int]int{
	'B': 0,
	'M': 1,
	'E': 2,
	'S': 3,
}

var BMESCount [4]float64

//BulidTransferProMaterix 	  求出
//TransferMatrix ：转移矩阵 4*4
//                    B   M   E  S  ALL
//                B   *   *   *  *  *    (取对数)
//                M   *   *   *  *  *
//                E   *   *   *  *  *
//                S   *   *   *  *  *
func BulidTransferProMaterix(path string) [SUM_STATUS][SUM_STATUS]float64 {

	var transferMaterix [SUM_STATUS][SUM_STATUS + 1]int64
	var transferProMaterix [SUM_STATUS][SUM_STATUS]float64

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(path, " is wrong")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	preStatus := 0
	for scanner.Scan() {
		s := scanner.Text()
		status := s[len(s)-1]
		if preStatus == 0 {
			preStatus = int(status)
			continue
		} else {
			transferMaterix[ma[preStatus]][ma[int(status)]]++
			preStatus = int(status)
		}
		transferMaterix[ma[int(status)]][SUM_STATUS]++
	}

	for i := 0; i < SUM_STATUS; i++ {
		for j := 0; j < SUM_STATUS; j++ {
			transferProMaterix[i][j] = math.Log(float64(transferMaterix[i][j]) / float64(transferMaterix[i][SUM_STATUS]))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Read error in BulidTransferProMaterix")
	}
	return transferProMaterix
}

//BulidEmitProMaterix 求出
//EmitProMaterix : 存储结构为四维数组 类型为map[rune]float64
//				   B   '汉字':probalitity(取对数)
//				   M
//                 E
//				   S
func BulidEmitProMaterix(path string) *[SUM_STATUS]map[rune]float64 {
	var ProMaterix = make(map[rune]*Feature)
	var EmitProMaterix [SUM_STATUS]map[rune]float64
	for i := 0; i < SUM_STATUS; i++ {
		EmitProMaterix[i] = make(map[rune]float64)
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(path, " is wrong")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		words := strings.Fields(s)
		if len(words) != 2 {
			continue
		}
		r := []rune(words[0])[0]
		if _, ok := ProMaterix[r]; !ok {
			status := int(words[1][0])
			ProMaterix[r] = &Feature{Count: 1}
			ProMaterix[r].BEMS[ma[status]]++
			BMESCount[ma[status]]++
		} else {
			status := int(words[1][0])
			ProMaterix[r].Count++
			ProMaterix[r].BEMS[ma[status]]++
			BMESCount[ma[status]]++
		}
	}

	for k := range ProMaterix {
		for i := 0; i < 4; i++ {
			EmitProMaterix[i][k] = math.Log(float64(ProMaterix[k].BEMS[i]+1) / float64(BMESCount[i]))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Read error in BulidEmitProMaterix")
	}
	return &EmitProMaterix
}

func HmmSaveTrainingFile() {
	TransferMatrix := BulidTransferProMaterix(seg.SegConfig["hmmTrainingFile"])
	EmitProMaterix := BulidEmitProMaterix(seg.SegConfig["hmmTrainingFile"])
	outFile, err := os.Create(seg.SegConfig["hmmModelFile"])

	defer outFile.Close()
	if err != nil {
		fmt.Println("Create Hmm Training File Failed")
	}
	outFile.WriteString("B           M                 E              S\n")

	outFile.WriteString("##prob_start\n")
	outFile.WriteString("-0.26268660809250016 -3.14e+100 -3.14e+100 -1.4652633398537678\n")

	outFile.WriteString("##TransferProMatrix\n")
	for i := 0; i < SUM_STATUS; i++ {
		for j := 0; j < SUM_STATUS; j++ {
			s := fmt.Sprintf("%f ", TransferMatrix[i][j])
			outFile.WriteString(s)
		}
		outFile.WriteString("\n")
	}

	outFile.WriteString("##EmitProMaterix\n")


	for i:=0;i<SUM_STATUS;i++{

		for k:=range  EmitProMaterix[i]{
			outFile.WriteString(string(k)+":")
			s:=fmt.Sprintf("%f",EmitProMaterix[i][k])
			outFile.WriteString(s+" ")
		}
		outFile.WriteString("\n")
	}
}
