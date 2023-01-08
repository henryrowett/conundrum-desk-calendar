package main

import (
	"fmt"
)

const ExitVal = -1024

type Node interface {
	Eval() float64
}

type LeafNode struct {
	Value float64
}

func (l LeafNode) Eval() float64 {
	return l.Value
}

type OperatorNode struct {
	Left, Right Node
	Operand     string
}

func (op OperatorNode) Operation() float64 {

	switch op.Operand {
	case "a":
		return op.Left.Eval() + op.Right.Eval()
	case "m":
		return op.Left.Eval() - op.Right.Eval()
	case "t":
		return op.Left.Eval() * op.Right.Eval()
	case "d":
		return op.Left.Eval() / op.Right.Eval()
	default:
		return -1000
	}
}

func (op OperatorNode) Eval() float64 {
	return op.Operation()
}

type NumCountdown struct {
	ExpTree OperatorNode
	Nums    []float64
}

func (t NumCountdown) IsEasy() bool {
	return false
}

func MakeRandomTree(p []float64) *NumCountdown {
	return &NumCountdown{
		ExpTree: OperatorNode{
			OperatorNode{LeafNode{p[0]}, LeafNode{p[1]}, "a"},
			OperatorNode{LeafNode{p[0]}, LeafNode{p[1]}, "a"},
			"t",
		},
		Nums: p,
	}
}

func TestVal(v float64) bool {
	if v != float64(int(v)) || v <= 0 {
		return false
	}
	return true
}

func GetTree(nums []float64) (NumCountdown, float64) {
	countdown := MakeRandomTree(nums)
	finalVal := countdown.ExpTree.Eval()

	if TestVal(finalVal) {
		return *countdown, finalVal
	}

	return GetTree(nums)
}

func main() {
	// seed := time.Now().UTC().UnixNano()
	// rng := rand.New(rand.NewSource(seed))

	tree, val := GetTree([]float64{1, 2})
	fmt.Printf("Tree is:\n %+v\n\n Final value is: %d", tree, int(val))
}

// if v, ok := TestVal(tree); ok {
//     return v
// }
