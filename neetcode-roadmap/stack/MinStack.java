import java.util.ArrayList;

// els is the stack where we store the value, 
// minVal is a "stack" that "track" all the 
// minimum value at a given index.
class MinStack {
    ArrayList<Integer> els;
    ArrayList<Integer> minVal;
    int min = Integer.MAX_VALUE;

    public MinStack() {
        els = new ArrayList<Integer>();
        minVal = new ArrayList<Integer>();
    }
    
    public void push(int val) {
        els.add(val);
        int curr = els.get(els.size()-1);
        if (curr < min){
            min = curr;
            minVal.add(min);
        } else{
            minVal.add(min);
        }
    }
    
    public void pop() {
        els.remove(els.size() - 1);
        minVal.remove(minVal.size() - 1);
        if (!minVal.isEmpty()) {
            min = minVal.get(minVal.size() - 1);
        // if the minVal array goes empty, we should 
        // restore the default value of min.
        } else {
            min = Integer.MAX_VALUE;
        }
    }
    
    public int top() {
        int res;
        res = els.get(els.size()-1);
        return res;
    }
    
    public int getMin() {
        return minVal.get(minVal.size()-1);
    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack obj = new MinStack();
 * obj.push(val);
 * obj.pop();
 * int param_3 = obj.top();
 * int param_4 = obj.getMin();
 */