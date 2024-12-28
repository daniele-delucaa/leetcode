# Brute force solution
# O(n * m) n = len(nums1), m = len(nums2)
class Solution:
    def nextGreaterElement(self, nums1: List[int], nums2: List[int]) -> List[int]:
        res = []
        index = 0
        for i, n1 in enumerate(nums1):
            j = nums2.index(n1)
            index = j 
            while index < len(nums2):
                if nums2[index] > n1:
                    res.append(nums2[index])
                    break
                if index == len(nums2)-1:
                    res.append(-1)
                index += 1
        return res
# Monotonic stack solution, using a hashmap where we save the
# value (as key) and indexes (as value) of an array. Iterate from
# nums2 and add n in stack if n is in nums1. If n is greater than
# the element on top of the stack, we found the next greater element
# of the element in the stack.
# So we pop from stack and get the index from hashmap and write the 
# result. 
# O(n + m)
    def nextGreaterElement(self, nums1: List[int], nums2: List[int]) -> List[int]:
        res = [-1] * len(nums1)
        stack = []
        occ = defaultdict(int)
        for x in nums1:
            occ[x] = nums1.index(x)
        for i, n in enumerate(nums2):
            while len(stack) > 0 and n > stack[len(stack)-1]:
                v = stack.pop()
                index = occ[v]
                res[index] = n
            if n in occ:
                stack.append(n)
        return res