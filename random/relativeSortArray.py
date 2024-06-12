from collections import defaultdict
 
class Solution:
    def relativeSortArray(self, arr1: List[int], arr2: List[int]) -> List[int]:
        result = []
        occurrences = defaultdict(int)
        for n in arr1:
            occurrences[n] += 1

        for n in arr2:
            if n in occurrences:
                # result.extend create a list, this list contains [n] multiplied for the occurrences of n, and this list get "appended" to "result"               
                result.extend([n] * occurrences[n])
                del occurrences[n]
        
        sorted_occ = sorted(occurrences.items())
        for k, v in sorted_occ:
            result.extend([k] * v)
        
        return result