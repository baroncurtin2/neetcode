class Solution:
    def longest_consecutive(self, nums: list[int]) -> int:
        num_set, longest = set(nums), 0

        for num in num_set:

            # if num - 1 is not in the set,
            # it means num is a starting point for the sequence
            if num - 1 not in num_set:
                count = 0

                while (num + count) in num_set:
                    count += 1
                longest = max(longest, count)
        return longest


def print_test(test_case: list[int], answer) -> None:
    print(
        f"Test Input: {test_case} \n"
        f"Answer: {answer} \n"
    )


if __name__ == "__main__":
    test1 = [100, 4, 200, 1, 3, 2]
    test2 = [0, 3, 7, 2, 5, 8, 4, 6, 0, 1]

    sol = Solution()
    longest_t1 = sol.longest_consecutive(test1)
    longest_t2 = sol.longest_consecutive(test2)

    print_test(test1, longest_t1)
    print_test(test2, longest_t2)
