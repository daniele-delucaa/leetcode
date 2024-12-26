# If it read an open bracket, push it in stack, when read 
# a closed bracket, check if it's the closed of the ones
# in top of the stack
class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        parentheses = {
            ")" : "(",
            "]" : "[",
            "}" : "{"
        }

        for ch in s:
            if ch == "(" or ch == "[" or ch == "{":
                stack.append(ch)
            else:
                if len(stack) == 0 or parentheses[ch] != stack[len(stack)-1]:
                    return False
                else:
                    stack.pop()
        if len(stack) == 0:
            return True
        else:
            return False
        