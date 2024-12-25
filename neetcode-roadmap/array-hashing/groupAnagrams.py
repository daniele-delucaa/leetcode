# First solution time limit exceeded 
# O(n) isAnagram 
# O(n * m^2) groupAnagram
class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        res: List[List[str]] = []
        #row = 0

        #strs_copy = strs.copy()
        for i in range(len(strs)):
            if strs[i] != "None":
                single_list = []
                single_list.append(strs[i])
                #res[row].append(strs[i])
                for j in range(i+1, len(strs)):
                    if self.isAnagram(strs[i], strs[j]):
                        #res[row].append(strs[j])
                        single_list.append(strs[j])
                        strs[j] = "None"
                #row += 1
                res.append(single_list)
            
        return res
    
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        # map that have arr (arr that count occurrences of characters in a string) as a key, and the lists of string associated as a value
        res = defaultdict(list) 

        for s in strs:
            # array of elements from a to z
            arr = [0] * 26 

            for ch in s:
                # count the occurrences of all characters in the string and increment the array (ord return the unicode value of ch)
                arr[ord(ch)-ord("a")] += 1
            res[tuple(arr)].append(s) # list can't be a key of a map, so we changed it to tuple

        return list(res.values())




    
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