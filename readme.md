# NLP工具库

## seg分词模块(MM,HMM，CRF)

- HMMmodel使用

```go
import (
	"fmt"

	"github.com/cocaer/goNLP/seg"
)
func main() {
	//seg.HmmSaveTraning()
	m := seg.NewHmmSeg() //基于hmm的优化实现
	fmt.Println(m.Cut("结婚的和尚未结婚的"))
}
```

``` txt
output:[结婚 的 和 尚未 结婚 的]
```


## 测试

Algorithm | Time | Precision | Recall|F-Measure
----|-------|------ | -----  |------|---------|
F-Measure | 3.21s 	   | 0.867 	   |0.896| 	0.881
ICTCLAS(2015版) | 0.55s   | 0.869   |0.914|0.891
jieba(C++版) | 0.26s   | 0.814   |0.809|0.811
THULAC_lite |0.62s|0.870|0.899|0.888
goNLP|2.01s|0.880|0.909|0.894|