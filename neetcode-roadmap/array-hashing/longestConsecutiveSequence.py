# Time limit excedeed
# convert nums in hash set, if n - 1 is not in set, n is 
# the start of a subsequence, looping through and build 
# the subsequence 
class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        num_set = set(nums)
        longest = 0 

        for n in nums:
            if n - 1 not in num_set:
                length = 1
                while n + length in num_set:   
                    length += 1
                longest = max(length, longest)

        return longest