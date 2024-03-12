# easy
class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        str = ""
        for element in s:
            i = 0
            for element2 in t:
                if element == element2:
                    str += element
                    t = t[i+1:]
                    break
                i += 1
        if str == s:
            return True
        return False       
    
# best
class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        if s == "":
            return True
        str = ""
        l = 0
        for element in t:
            if l == len(s): 
                break
            if element == s[l]:
                str += element
                l += 1
        if str == s:
            return True
        return False