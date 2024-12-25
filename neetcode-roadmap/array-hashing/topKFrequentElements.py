# HashMap solution
class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        res = []
        freq = defaultdict(int)
        for x in nums:
            freq[x] += 1

        elem = 0
        for i in range(k):
            max = 0
            for k, v in freq.items():
                if v > max:
                    max = v
                    elem = k 
            res.append(elem)
            del freq[elem]
        return res
    

    # Bucket sort solution
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        res = []
        count = {}
        freq = [[] for i in range(len(nums) + 1)]

        for x in nums:
            count[x] = count.get(x, 0) + 1

        for key, v in count.items():
            freq[v].append(key)

        # looping through array from last to first element
        for i in range(len(freq)-1, 0, -1):
            for n in freq[i]:
                res.append(n)
                if len(res) == k:
                    return res