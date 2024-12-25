class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        d1 = {}
    
        for x in s:
            if x in d1:
                d1[x] += 1
            else:
                d1[x] = 1

        for y in t:
            if y not in d1:
                return False
            elif d1[y] == 1:
                del d1[y]
            # if y is in dictionary d1 and occurrences are more than 1
            else:
                d1[y] -= 1
        
        # if d1 is empty, the strings are anagram and it will return true, else if it is not empty will return false
        return not d1