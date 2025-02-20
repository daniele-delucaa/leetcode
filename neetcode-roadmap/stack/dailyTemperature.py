class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        res = [0] * len(temperatures)
        # dict with index as a key and temperature as value
        d = defaultdict(int)
        stack = []
        for i, x in enumerate(temperatures):
            while stack and x > d[stack[-1]]:
                idx = stack.pop()
                res[idx] = i - idx
            d[i] = x
            stack.append(i)
        return res
                
                 
