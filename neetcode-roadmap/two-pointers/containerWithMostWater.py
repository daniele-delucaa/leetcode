# Brute force solution, time limit exceeded
class Solution:
    def maxArea(self, height: List[int]) -> int:
        res = 0
        i = 0
        while i < len(height):
            j = i+1
            while j < len(height):
                if i == j:
                    j += 1
                    continue
                width = min(height[i], height[j])
                length = j - i
                v = width * length 
                if v > res:
                    res = v
                j += 1
            i += 1
        return res
    
# Two pointers, set one pointer to the left and one
# to the right of the array. Always move the pointer 
# that points to the lower line.
    def maxArea(self, height: List[int]) -> int:
        res = 0
        l, r = 0, len(height)-1
        i = 0
        while i < len(height):
            width = min(height[l], height[r])
            length = r - l
            v = width * length
            if v > res:
               res = v
            if width == height[l]:
               l += 1
            elif width == height[r]:
               r -= 1
            i += 1
        return res
                