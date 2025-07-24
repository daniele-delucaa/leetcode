import java.util.HashSet;

public class LongestSubstringWithoutRepeatingCharacters {
    public int lengthOfLongestSubstring(String s) {
        int left = 0;
        int right = 0;
        int res = 0;
        HashSet<Character> occurrences = new HashSet<>();

        while (right < s.length()) {
            if (!occurrences.contains(s.charAt(right))) {
                occurrences.add(s.charAt(right));
                res = Math.max(res, right - left + 1); 
                right++;
            } else {
                occurrences.remove(s.charAt(left));
                left++;
            }
        }
        return res;
    }
}
