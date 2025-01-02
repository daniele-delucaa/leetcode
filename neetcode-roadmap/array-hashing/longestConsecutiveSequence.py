# convert nums in hash set, if n - 1 is not in set, n is 
# the start of a subsequence, looping through (while n + 1 is in set) 
# and build the subsequence 
class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        nums = set(nums)
        longest = 0 

        for n in nums:
            if n - 1 not in nums:
                length = 1
                while n + length in nums:   
                    length += 1
                longest = max(length, longest)

        return longest