import java.util.Stack;

class Solution {
    public int[] finalPrices(int[] prices) {
        int n = prices.length;
        int[] res = new int[n];
        Stack<Integer> stack = new Stack<>();

        for (int i = 0; i < n; i++) {
            while (!stack.isEmpty() && prices[stack.peek()] >= prices[i]) {
                int idx = stack.pop();
                res[idx] = prices[idx] - prices[i];
            }
            // saving index
            stack.push(i); 
        }

        // the prices without discount remain the same
        while (!stack.isEmpty()) {
            int idx = stack.pop();
            res[idx] = prices[idx];
        }

        return res;
    }
}