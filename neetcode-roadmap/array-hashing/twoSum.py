class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        d = defaultdict(int)
        res = []
        for i, n in enumerate(nums):
            y = target - n
            if y in d:
                res.append(d[y])
                res.append(i)
                return res
            d[n] = i