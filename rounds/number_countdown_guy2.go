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

// TODO: Impliment
func (t NumCountdown) IsEasy() bool {
	return false
}

func generateLeafs(nums []float64, leaves ...*LeafNode) {
	for i, n := range nums {
		leaves[i] = &LeafNode{
			Value: n,
		}
	}
}

// TODO: Impliment
func makeRandomTree(nums []float64) *NumCountdown {

	var l1, l2, l3, l4, l5, l6 *LeafNode
	generateLeafs(nums, l1, l2, l3, l4, l5, l6)

	return &NumCountdown{
		ExpTree: OperatorNode{
			OperatorNode{LeafNode{nums[0]}, LeafNode{nums[1]}, "a"},
			OperatorNode{LeafNode{nums[0]}, LeafNode{nums[1]}, "a"},
			"t",
		},
		Nums: nums,
	}
}

func testVal(c *NumCountdown) (float64, bool) {
	v := c.ExpTree.Eval()
	if v != float64(int(v)) || v <= 0 {
		return v, false
	}
	return v, true
}

func getTree(nums []float64) (NumCountdown, float64) {
	countdown := makeRandomTree(nums)

	if v, ok := testVal(countdown); ok {
		return *countdown, v
	}

	return getTree(nums)
}

func main() {
	// seed := time.Now().UTC().UnixNano()
	// rng := rand.New(rand.NewSource(seed))

	numCountdown, val := getTree([]float64{1, 2})
	fmt.Printf("Tree is:\n %+v\n\n Final value is: %d", numCountdown, int(val))
}

// // Generator function (not used)
// func GenerateLeafs(nums []float64) chan *LeafNode {
// 	c := make(chan *LeafNode)
// 	go func() {
// 		for _, i := range nums {
// 			c <- &LeafNode{
// 				Value: i,
// 			}
// 		}
// 		close(c)
// 	}()
//
// }
