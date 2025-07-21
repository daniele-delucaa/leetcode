class BestTimeToBuyAndSellStock {
    public int maxProfit(int[] prices) {
        int maxProfit = 0;
        int left = 0;
        int right = 1;

        while (right < prices.length){
            int profit;
            if (prices[left] < prices[right]){
                profit = prices[right] - prices[left];
                maxProfit = Math.max(maxProfit, profit);
            }else {
                left = right;
            }
            right++;
        }
        return maxProfit;
    }
}