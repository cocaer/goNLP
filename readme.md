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


## TODO
