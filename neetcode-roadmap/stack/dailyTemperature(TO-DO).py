class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        res = [0] * len(temperatures)
        stack = []
        stack.append(temperatures[0])
        for i, n in enumerate(temperatures):
            if i == 0:
                continue
            if temperatures[i] > stack[len(stack)-1]:
                 
