package main

import (
	"fmt"
	"math/rand"
	"time"
)

const ExitVal = -1024

var Operands = [...]string{"a", "m", "d", "t"}
var Sides = [...]string{"l", "r"}

func testVal(v float64) (float64, bool) {
	if v != float64(int(v)) || v <= 0 {
		return v, false
	}
	return v, true
}

type Node interface {
	Eval() float64
	AddNode(n *OperatorNode, s string, rng *rand.Rand)
	NoLeaves() int
}

type LeafNode struct {
	Value float64
}

func (l LeafNode) Eval() float64 {
	return l.Value
}

func (l LeafNode) NoLevaes() int {
	return 1
}

// Hella ugly need to fix
func (l LeafNode) AddNode(n *OperatorNode, s string, rng *rand.Rand) {
	return
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

// TODO: want to fail if this operation does not return an integer value above zero
func (op OperatorNode) Eval() float64 {
	return op.Operation()
}

func (op OperatorNode) NoLeaves() int {
	n := 0

	if op.Right == nil {
		n += 1
	} else {
		n += op.Right.NoLeaves()
	}
	if op.Left == nil {
		n += 1
	} else {
		n += op.Left.NoLeaves()
	}

	return n
}

// HELPPPPP
func (op *OperatorNode) AddNode(n *Node, s string, rng *rand.Rand) {
	var side string

	if s == "r" {
		if op.Right == nil {
			op.Right = n
		} else {
			side = Sides[rng.Intn(len(Sides))]
			op.Right.AddNode(n, side, rng)
		}
	} else {
		if op.Left == nil {
			op.Left = n
		} else {
			side = Sides[rng.Intn(len(Sides))]
			op.Left.AddNode(n, side, rng)
		}
	}
}

func generateLeafs(nums []float64, leaves ...*LeafNode) {
	for i, n := range nums {
		leaves[i] = &LeafNode{
			Value: n,
		}
	}
}

// Always 5 branches (including root), need to combine them diffierently
func generateTreeStruct(rng *rand.Rand) OperatorNode {
	possTree := OperatorNode{
		Left:    nil,
		Right:   nil,
		Operand: Operands[rand.Intn(len(Operands)-1)],
	}

	for possTree.NoLeaves() < 6 {
		fmt.Printf("No. leaves: %d\n", int(possTree.NoLeaves()))
		side := Sides[rand.Intn(len(Sides))]
		possTree.AddNode(
			&OperatorNode{
				Left:    nil,
				Right:   nil,
				Operand: Operands[rand.Intn(len(Operands))],
			},
			side,
			rng,
		)
	}
	return possTree
}

type NumCountdown struct {
	ExpTree OperatorNode
	Nums    []float64
}

// TODO: Impliment
func (t NumCountdown) IsEasy() bool {
	return false
}

// TODO: Impliment
func makeRandomTree(nums []float64, rng *rand.Rand) *NumCountdown {

	var l1, l2, l3, l4, l5, l6 *LeafNode
	generateLeafs(nums, l1, l2, l3, l4, l5, l6)
	treeStruct := generateTreeStruct(rng)

	for _, l := range []*LeafNode{l1, l2, l3, l4, l5, l6} {
		side := Sides[rand.Intn(len(Sides))]
		treeStruct.AddNode(l, side, rng)
	}

	if testVal(treeStruct.Eval()) {
		return &NumCountdown{
			ExpTree: treeStruct,
			Nums:    nums,
		}
	} else {
		makeRandomTree(nums, rng)
	}
}

func getTree(nums []float64) (NumCountdown, float64) {
	countdown := makeRandomTree(nums)
	val := countdown.ExpTree.Eval()

	if v, ok := testVal(val); ok {
		return *countdown, v
	}

	return getTree(nums)
}

func main() {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	// numCountdown, val := getTree([]float64{1, 2})
	// fmt.Printf("Tree is:\n %+v\n\n Final value is: %d", numCountdown, int(val))

	countdown := makeRandomTree([]float64{1, 2, 3, 4, 5, 6}, rng)

	fmt.Printf("%+v", countdown)

	fmt.Println(countdown.Expr.Eval())

}

// // Generator function for operators
// func generateOperands() chan string {
// 	c := make(chan string)
// 	go func() {
// 		for i := 0; i < 100; i++ {
// 			idx := rand.Intn(len(Operands))
// 			c <- Operands[idx]
// 		}
// 		close(c)
// 	}()
// 	return c
// }
