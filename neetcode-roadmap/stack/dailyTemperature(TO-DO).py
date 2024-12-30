class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        res = [0] * len(temperatures)
        stack = []
        i = 1
        stack.append(temperatures[0])
        while i < len(temperatures):
            if temperatures[i] > stack[len(stack)-1]:
                res[i] = i - temperatures.index(stack[len(stack)-1])
                stack.pop()
                stack.append(temperatures[i])
            else:
                stack.append(temperatures[i])
            i += 1
        return res
                
                 
