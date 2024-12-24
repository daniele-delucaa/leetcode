class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        dict = {}
        for x in nums:
            if x in dict:
                dict[x] += 1
            else:
                dict[x] = 1

        for y in dict:
            if dict[y] > 1:
                return True
            
        return False
