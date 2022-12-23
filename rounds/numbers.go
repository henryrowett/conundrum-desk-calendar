package rounds

import (
	"countdown/utils"
	"fmt"
	_ "go/parser"
	"math/rand"
	"strings"
)

type Numbers struct{}

func (n *Numbers) Generate() Result {
	// generate a pseudo-random seed
	rand.Seed(int64(timeNowF.YearDay()))

	// fetch 2 big and 4 small numbers
	var numbers []int
	for i := 1; i <= 4; i++ {
		numbers = append(numbers, utils.NumberSmallSet[rand.Intn(len(utils.NumberSmallSet))])
	}

	for i := 1; i <= 2; i++ {
		numbers = append(numbers, utils.NumberLargeSet[rand.Intn(len(utils.NumberLargeSet))])
	}

	// randomly arrange the numbers
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	// 5 random operators
	var operators []string
	for i := 1; i <= 5; i++ {
		operators = append(operators, utils.NumberOperators[rand.Intn(len(utils.NumberOperators))])
	}

	// use a variation of polish notation to apply operators
	// https://en.wikipedia.org/wiki/Polish_notation
	// we take the first two number elements and first operator
	// and calculate the result and append that to then be operated on
	today, ops := recurse(operators, numbers, []string{})
	yesterday := fmt.Sprintf("%s %s", strings.Join(ops, ","), strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers)), ","), "[]"))

	// re-randomise numbers before sending
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	// return result
	return Result{
		Round:     numbersRound,
		Today:     fmt.Sprintf("%s = %d", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers)), ","), "[]"), today),
		Yesterday: yesterday,
	}
}

func recurse(operators []string, numbers []int, ops []string) (int, []string) {

	if len(operators) == 0 {
		return numbers[0], ops
	}

	op := operators[0]
	a := numbers[0]
	b := numbers[1]

	switch op {
	case "+":
		ops = append(ops, op)
		z := a + b
		n := numbers[2:]
		r := append([]int{z}, n...)
		return recurse(operators[1:], r, ops)
	case "-":
		ops = append(ops, op)
		z := a - b
		n := numbers[2:]
		r := append([]int{z}, n...)
		return recurse(operators[1:], r, ops)
	case "*":
		ops = append(ops, op)
		z := a * b
		n := numbers[2:]
		r := append([]int{z}, n...)
		return recurse(operators[1:], r, ops)
	case "/":
		var z int
		if a%b != 0 {
			z = a / b
			ops = append(ops, op)
		} else if b%a != 0 {
			z = b % a
			ops = append(ops, op)
		} else {
			z = a + b
			ops = append(ops, "+")
		}

		n := numbers[2:]
		r := append([]int{z}, n...)
		return recurse(operators[1:], r, ops)
	}

	return 0, []string{}
}
