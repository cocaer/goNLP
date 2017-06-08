package hmm

import "os"
import "fmt"
import "bufio"
import "strings"

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

//BulidTransferProMaterix 	  求出
//TransferMatrix ：转移矩阵 4*4
//                    B   M   E  S  ALL
//                B   *   *   *  *  *
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
			transferProMaterix[i][j] = float64(transferMaterix[i][j]) / float64(transferMaterix[i][SUM_STATUS])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Read error in BulidTransferProMaterix")
	}
	return transferProMaterix
}

//BulidEmitProMaterix 求出
//EmitProMaterix : 存储结构为map
//				   key:汉字
//				   value:Feature结构体
func BulidEmitProMaterix(path string) map[rune]*Feature {
	var EmitPropMaterix = make(map[rune]*Feature)

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
		if _, ok := EmitPropMaterix[r]; !ok {
			EmitPropMaterix[r] = &Feature{Count: 1}
		} else {
			status := int(words[1][0])
			EmitPropMaterix[r].Count++
			EmitPropMaterix[r].BEMS[ma[status]]++
		}
	}

	for k := range EmitPropMaterix {
		for i := 0; i < 4; i++ {
			EmitPropMaterix[k].BEMSPro[i] = float64(EmitPropMaterix[k].BEMS[i]) / float64(EmitPropMaterix[k].Count)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Read error in BulidEmitProMaterix")
	}
	return EmitPropMaterix
}
