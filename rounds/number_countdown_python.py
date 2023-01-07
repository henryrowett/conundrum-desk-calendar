import numpy as np
import numpy.typing as npt
from typing import List, Generator, Tuple, Dict, Callable
import operator
import random
import copy

ops: Dict[Callable, str] = {
        operator.add: "+",
        operator.sub: "-",
        operator.mul: "x",
        operator.truediv: "/"
        }

def next_op() -> Generator:
    for i in range(1000):
        op = random.choice(list(ops.keys()))
        yield op


s_nums = np.arange(1, 9)
l_nums = np.array([25, 50, 75, 100])


def get_nums() -> npt.NDArray:
    nums = list(np.random.choice(s_nums, 4, replace=False)) 
    nums.extend(np.random.choice(l_nums, 2, replace=False)) 
    np.random.shuffle(nums)
    return nums

def answer(nums: npt.NDArray) -> Tuple[int, List[str]]:
    ops_carried = []
    temp_nums = copy.copy(nums)
    
    while len(temp_nums) > 1:
        a, b = np.random.choice(temp_nums, 2, replace=False)
        op = next(next_op())
        new_num = op(a, b)
        while new_num <= 0 or new_num != int(new_num):
            a, b = np.random.choice(temp_nums, 2, replace=False)
            op = next(next_op())
            new_num = op(a, b)

        ops_carried.append(f"{a} {ops[op]} {b} = {new_num}")

        temp_nums.remove(a)
        temp_nums.remove(b)
        temp_nums.append(new_num)

    return temp_nums[0], ops_carried

def answer2(nums: npt.NDArray) -> Tuple[int, List[str]]:
    ans, ops = answer(nums)
    if ans < 100 or ans > 999:
        ans, ops = answer2(nums)

    return ans, ops

def validate_hardness(nums: npt.NDArray, ans: int) -> int:
    ...

def main() -> None:
    nums = get_nums()
    ans, operations_used = answer2(nums)

    print(f"Numbers are {nums}")
    print(f"Number to get to is {ans}")
    print("A method is:")
    print("\n".join(operations_used))


main()
