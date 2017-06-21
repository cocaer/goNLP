package data

//
//bhmm(better hmm)不仅依赖正确的bmes序列，还依赖
//rmm 和 mm产生的bmes序列
//
import "os"
import "bufio"
import "strings"

func Merge(in1, in2, in3, out string) {

	ina, _ := os.Open(in1)
	inb, _ := os.Open(in2)
	inc, _ := os.Open(in3)
	outa, _ := os.Create(out)
	defer ina.Close()
	defer inb.Close()
	defer inc.Close()

	sa := bufio.NewScanner(ina)
	sb := bufio.NewScanner(inb)
	sc := bufio.NewScanner(inc)

	for sa.Scan() && sb.Scan() && sc.Scan() {
		hanzi := strings.Fields(sa.Text())[0]
		s1 := strings.Fields(sa.Text())[1]
		s2 := strings.Fields(sb.Text())[1]
		s3 := strings.Fields(sc.Text())[1]
		outa.WriteString(hanzi + " " + s1 + " " + s2 + " " + s3 + "\n")
	}
}
