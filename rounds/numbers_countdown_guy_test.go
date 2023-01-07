package rounds

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestGuyGo(t *testing.T) {
	seed := rand.NewSource(int64(timeNowF.YearDay()))
	rng := rand.New(seed)

	nums := getNumbers(rng)
	fmt.Println(nums)

	a, o := answer2(nums, rng)

	fmt.Println(a)
	fmt.Println(o)
}
