from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        seen = {}

        for i, x in enumerate(nums):
            x_target = target - x

            if x_target not in seen:
                seen[x] = i
            else:
                return [seen[x_target], i]
