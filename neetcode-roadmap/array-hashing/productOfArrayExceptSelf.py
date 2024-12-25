# Brute force solution time limit excedeed
class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        supp = []
        #res = []
        #for n in nums:
        #    supp.append(n)
 
        for i in range(len(nums)):
            mul = 1
            for j in range(len(nums)):
                if i != j:
                    mul *= nums[j] 
            supp.append(mul)
            mul = 1
        return supp
    
    # for any nums[i], calculate its left product and its right product
    # without including nums[i]. Then multiply left and right product. 
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        pref = []
        suff = [0] * len(nums)
        
        mul = 1
        pref.append(1)
        # prefix products
        for i in range(1, len(nums)):
            mul *= nums[i-1]
            pref.append(mul)

       
        mul = 1
        # suffix products
        for i in range(len(nums)-2, -1, -1):
            mul *= nums[i+1]
            suff[i] = mul
        suff[len(nums)-1] = 1

        for i in range(len(nums)):
            nums[i] = pref[i] * suff[i]
        return nums