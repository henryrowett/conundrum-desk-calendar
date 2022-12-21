package rounds

type Numbers struct{}

func (n *Numbers) Generate() Result {
	// randomly Generate 6 Numbers from the set
	// - [0,1,2,3,4,5,6,7,8,9,10]
	// - [25,50,75,100]

	// perform a mathematical operation using these numbers

	// return result
	return Result{
		Round:     numbersRound,
		Today:     "[] = 12",
		Yesterday: "1*2*3*4*5*6",
	}
}
