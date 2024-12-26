# Make two array sum_left, sum_right where sumLeft[i]
# is the sum of all the numbers to the left of index i
# and where sumRight[i] is the sum of all the numbers to 
# the right of index i
class Solution:
    def pivotIndex(self, nums: List[int]) -> int:
        res = -1
        sum_left = [0] * len(nums)
        sum_right = [0] * len(nums)
        counter_left = 0
        counter_right = 0
        for index, n in enumerate(nums):
            sum_left[index] = counter_left
            counter_left += n

        for index in range(len(nums)-1, -1, -1):
            sum_right[index] = counter_right
            counter_right += nums[index]

        for index, _ in enumerate(sum_left):
            if sum_left[index] == sum_right[index]:
                return index

        return res  
