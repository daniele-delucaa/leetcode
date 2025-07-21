class sortColors {

    // Counting sort solution 
    // Not in place, in-place required (use Dutch national flag's algorithm)
    // Time complexity: O(n)
    // Space complexity: O(n)
    public void sortColors(int[] nums) {
        int max = 0;
        for (int i = 0; i < nums.length; i++){
            if (nums[i] > max){
                max = nums[i];
            }
        }
        int[] supp = new int[max+1];
        for (int i = 0; i < nums.length; i++){
            supp[nums[i]]++;
        }
        int index = 0;
        for (int suppIndex = 0; suppIndex <= max; suppIndex++) {
            for (int j = 0; j < supp[suppIndex]; j++) {
                nums[index] = suppIndex;
                index++;
            }
        }
    }
}
