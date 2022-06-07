from typing import List


class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        n = len(nums)
        result = [1] * n

        right_pass = 1
        for i in range(n):
            result[i] = right_pass
            right_pass *= nums[i]
            print(right_pass, result, nums[i])

        left_pass = 1
        for i in range(n - 1, -1, -1):
            result[i] *= left_pass
            left_pass *= nums[i]
            print(left_pass, result, nums[i])

        return result


if __name__ == "__main__":
    a = [1, 2, 3, 4]
    s = Solution()
    print(s.productExceptSelf(a))
