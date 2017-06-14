package hmm

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	seg "github.com/cocaer/goNLP/seg/config"
)

const IMPOSSIBLEPRO = -3.14e+10

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

var am = map[int]rune{
	0: 'B',
	1: 'M',
	2: 'E',
	3: 'S',
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
	preStatus := -1
	for scanner.Scan() {
		s := scanner.Text()
		status := s[len(s)-1]
		transferMaterix[ma[int(status)]][SUM_STATUS]++
		if preStatus == -1 {
			preStatus = int(status)
			continue
		} else {
			transferMaterix[ma[preStatus]][ma[int(status)]]++
			preStatus = int(status)
		}
	}

	for i := 0; i < SUM_STATUS; i++ {
		for j := 0; j < SUM_STATUS; j++ {

			if transferMaterix[i][j] == 0 {
				transferProMaterix[i][j] = IMPOSSIBLEPRO
				continue
			}
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
			EmitProMaterix[i][k] = math.Log((float64(ProMaterix[k].BEMS[i] + 1)) / float64(BMESCount[i]))
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

	outFile.WriteString("package data \n" +
		"const SUM_STATUS  =4\n" +
		"var StartProMaterix =[SUM_STATUS]float64{" +
		"-0.26268660809250016, -3.14e+10 ,-3.14e+10, -1.4652633398537678" +
		"}\n")

	outFile.WriteString("var TransferMatrix  =[SUM_STATUS][SUM_STATUS]float64{")

	for i := 0; i < SUM_STATUS; i++ {
		var s = fmt.Sprintf("{%f,%f,%f,%f}", TransferMatrix[i][0],
			TransferMatrix[i][1],
			TransferMatrix[i][2],
			TransferMatrix[i][3])
		outFile.WriteString(s)
		outFile.WriteString(",")
	}
	outFile.WriteString("}\n")

	outFile.WriteString("var EmitProMaterix  =" +
		"&[SUM_STATUS]map[rune]float64{")

	for i := 0; i < SUM_STATUS; i++ {
		outFile.WriteString("{\n")
		for k := range EmitProMaterix[i] {
			s := fmt.Sprintf("'%c':%f,\n", k, EmitProMaterix[i][k])
			outFile.WriteString(s)
		}
		outFile.WriteString("},\n")
	}
	outFile.WriteString("}")

}
