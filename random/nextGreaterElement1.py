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
                