package hmm

const (
	B = iota
	E
	M
	S
	SUM_STATUS
)

type HMMModel struct {
	InitStatus     [SUM_STATUS]float64
	TransferMatrix [SUM_STATUS][SUM_STATUS]float64
	EmitMatrix     [SUM_STATUS][SUM_STATUS]float64

	EmitProbB map[rune]float64
	EmitProbE map[rune]float64
	EmitProbM map[rune]float64
	EmitProbS map[rune]float64
}
