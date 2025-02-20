import java.util.Stack;

class Solution {
    public String removeDuplicates(String s) {
        Stack<String> stack = new Stack<>();
        for (int i = 0; i < s.length(); i++){
            if (!stack.empty() && stack.peek().equals(s.substring(i, i+1))){
                stack.pop();
            }else {
                stack.push(s.substring(i, i+1));
            }
        }
        StringBuilder sb = new StringBuilder();
        for(String ch : stack){
            sb.append(ch);
        }
        return sb.toString();
    }
}