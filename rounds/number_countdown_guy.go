package rounds

import (
	"fmt"
	"math/rand"
)

func myCopy(a []int) []int {
	var b []int
	for i := 0; i < len(a); i++ {
		b = append(b, a[i])
	}
	return b
}

func split(a, b []int) (int, int, int, int) {
	return a[0], a[1], b[0], b[1]
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func randChoice(arr []int, n int, rng *rand.Rand) ([]int, []int) {
	var chosenElems []int
	var idx []int
	for i := 0; i < n; i++ {
		j := 0
		if len(arr) > 1 {
			j = rng.Intn(len(arr) - 1)
		}
		chosenElems = append(chosenElems, arr[j])
		idx = append(idx, j)
		arr = remove(arr, j)
	}
	return chosenElems, idx
}

func ops(a, b float64, op string) float64 {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "/":
		return a / b
	case "x":
		return a * b
	default: // Never reached
		return a
	}
}

func getOperation(i int) string {
	ops := []string{"+", "-", "/", "x"}
	return ops[i]
}

func generateOperation(rng *rand.Rand) string {
	i := rng.Intn(4)
	return getOperation(i)
}

func getNumbers(rng *rand.Rand) []int {
	small := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	large := []int{25, 50, 75, 100}

	var nums []int
	ss, _ := randChoice(small, 4, rng)
	nums = append(nums, ss...)
	ls, _ := randChoice(large, 2, rng)
	nums = append(nums, ls...)

	shuffled := nums
	rand.Shuffle(len(nums), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}

// TODO: get repeated access to elements in the array, something in the random choice
func answer(nums []int, rng *rand.Rand) (int, []string) {
	var opsCarried []string
	tempNums := myCopy(nums)

	for len(tempNums) > 1 {
		a, b, i, j := split(randChoice(tempNums, 2, rng))
		op := generateOperation(rng)
		new_num := ops(float64(a), float64(b), op)

		for (new_num <= 0) || (new_num != float64(int(new_num))) {
			a, b, i, j = split(randChoice(tempNums, 2, rng))
			op = generateOperation(rng)
			new_num = ops(float64(a), float64(b), op)
		}

		s := fmt.Sprintf("%d", int(a)) + op + fmt.Sprintf("%d", int(b)) + " = " + fmt.Sprintf("%d", int(new_num))
		opsCarried = append(opsCarried, s)

		tempNums = remove(tempNums, i)
		tempNums = remove(tempNums, j)
		tempNums = append(tempNums, int(new_num))
	}

	return tempNums[0], opsCarried
}

// On each loop of while, the nums that is put in has already been changed in the previous iteration?
// TODO: Find a way to pass nums as a myCopy or make a myCopy in the answer function so that the originol nums are not altered
func answer2(nums []int, rng *rand.Rand) (int, []string) {
	fmt.Println("nums: ", nums)
	ans, ops := answer(nums, rng)
	for (ans < 100) || (ans > 999) {
		fmt.Println("nums: ", nums)
		ans, ops = answer(nums, rng)
	}

	return ans, ops
}
