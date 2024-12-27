class Solution:
    def clearDigits(self, s: str) -> str:
        res = ""
        stack = []
        for ch in s:
            if ch.isdigit():
                stack.pop()
            else:
                stack.append(ch)
        
        for ch in stack:
            res += ch
             
        return res