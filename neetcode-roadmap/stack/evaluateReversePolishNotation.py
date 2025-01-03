# if token[i] is a number, push it into the stack, if it is
# an "oparator", pop the last two numbers and do the operations, 
# lastly push the result of the operation in to the stack
class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        number_stack = []
        for ch in tokens:
            if ch == "+":
                number_stack.append(number_stack.pop() + number_stack.pop())
            elif ch == "-":
                n1, n2 = number_stack.pop(), number_stack.pop()
                number_stack.append(n2-n1)
            elif ch == "*":
                number_stack.append(number_stack.pop() * number_stack.pop())
            elif ch == "/":
                n1, n2 = number_stack.pop(), number_stack.pop()
                number_stack.append(int(n2 / n1))
            else:
                number_stack.append(int(ch))
        return number_stack[0]
