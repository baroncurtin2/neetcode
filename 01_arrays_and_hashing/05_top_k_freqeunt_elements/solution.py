from typing import List
from collections import defaultdict, Counter


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        freq = defaultdict(list)

        for key, ct in Counter(nums).items():
            freq[ct].append(key)

        res = []
        for times in reversed(range(len(nums) + 1)):
            res.extend(freq[times])

            if len(res) >= k:
                return res[:k]

        return res[:k]
