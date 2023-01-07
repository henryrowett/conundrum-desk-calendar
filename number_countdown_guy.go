package main

import (
	"fmt"
	"math/rand"
	"time"
)

func my_copy(arr []int) []int {
	var new_arr []int
	for i := 0; i < len(arr); i++ {
		new_arr = append(new_arr, arr[i])
	}
	return new_arr
}

func split_elems(a, b []int) (int, int, int, int) {
	return a[0], a[1], b[0], b[1]
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func rand_choice(arr []int, n int, rng *rand.Rand) ([]int, []int) {
	var chosen_elems []int
	var idx []int
	for i := 0; i < n; i++ {
		j := 0
		if len(arr) > 1 {
			j = rng.Intn(len(arr) - 1)
		}
		chosen_elems = append(chosen_elems, arr[j])
		idx = append(idx, j)
		arr = remove(arr, j)
	}
	return chosen_elems, idx
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

func get_oper(i int) string {
	pos_ops := []string{"+", "-", "/", "x"}
	return pos_ops[i]
}

func oper_gen(rng *rand.Rand) string {
	i := rng.Intn(4)
	return get_oper(i)
}

func get_nums(rng *rand.Rand) []int {
	s_nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	l_nums := []int{25, 50, 75, 100}

	var nums []int
	ss, _ := rand_choice(s_nums, 4, rng)
	nums = append(nums, ss...)
	ls, _ := rand_choice(l_nums, 2, rng)
	nums = append(nums, ls...)

	shuf_nums := []int(nums)
	rand.Shuffle(len(nums), func(i, j int) {
		shuf_nums[i], shuf_nums[j] = shuf_nums[j], shuf_nums[i]
	})

	return shuf_nums
}

// TODO: get repeated access to elements in the array, something in the random choice
func answer(nums []int, rng *rand.Rand) (int, []string) {
	var ops_carried []string
	temp_nums := my_copy(nums)

	for len(temp_nums) > 1 {
		a, b, i, j := split_elems(rand_choice(temp_nums, 2, rng))
		op := oper_gen(rng)
		new_num := ops(float64(a), float64(b), op)

		for (new_num <= 0) || (new_num != float64(int(new_num))) {
			a, b, i, j = split_elems(rand_choice(temp_nums, 2, rng))
			op = oper_gen(rng)
			new_num = ops(float64(a), float64(b), op)
		}

		s := fmt.Sprintf("%d", int(a)) + op + fmt.Sprintf("%d", int(b)) + " = " + fmt.Sprintf("%d", int(new_num))
		ops_carried = append(ops_carried, s)

		temp_nums = remove(temp_nums, i)
		temp_nums = remove(temp_nums, j)
		temp_nums = append(temp_nums, int(new_num))
	}

	return temp_nums[0], ops_carried
}

// On each loop of while, the nums that is put in has already been changed in the previous iteration?
// TODO: Find a way to pass nums as a copy or make a copy in the answer function so that the originol nums are not altered
func answer2(nums []int, rng *rand.Rand) (int, []string) {
	fmt.Println("nums: ", nums)
	ans, ops := answer(nums, rng)
	for (ans < 100) || (ans > 999) {
		fmt.Println("nums: ", nums)
		ans, ops = answer(nums, rng)
	}

	return ans, ops
}

func main() {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	nums := get_nums(rng)
	fmt.Println(nums)

	a, o := answer2(nums, rng)

	fmt.Println(a)
	fmt.Println(o)
}
