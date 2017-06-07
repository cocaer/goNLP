package hmm

import "os"
import "fmt"
import "bufio"

//Feature 一个汉字一个Feature
type Feature struct {
	Count   int
	BEMS    [SUM_STATUS]int
	BEMSPro [SUM_STATUS]float64
}

//BulidMaterix 求出
//TransferMatrix ：转移矩阵 4*4
//				      B   M   E  S  ALL
//				  B   *   *   *  *  *
//				  M   *   *   *  *  *
//			      E   *   *   *  *  *
//				  S   *   *   *  *  *
func BulidTransferMaterix(path string) {

	var transferMaterix [SUM_STATUS][SUM_STATUS + 1]int64
	var transferProMaterix [SUM_STATUS][SUM_STATUS]float64
	ma := map[int]int{
		'B': 0,
		'M': 1,
		'E': 2,
		'S': 3,
	}

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

	for i := 0; i < 4; i++ {
		fmt.Println(transferProMaterix[i])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Read error in HmmTraining")
	}

}
