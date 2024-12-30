import java.util.ArrayList;

class MyQueue {
    ArrayList<Integer> q;
    public MyQueue() {
        q = new ArrayList<Integer>();
    }
    
    public void push(int x) {
        q.add(x);
    }
    
    public int pop() {
        int res = q.get(0);
        q.remove(0);
        /*for (int i = 1; i < q.size()-1; i++){
            q.set(i-1, q.get(i));
        }*/
        return res;
    }
    
    public int peek() {
        return q.get(0);
    }
    
    public boolean empty() {
        return q.isEmpty();
    }
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * MyQueue obj = new MyQueue();
 * obj.push(x);
 * int param_2 = obj.pop();
 * int param_3 = obj.peek();
 * boolean param_4 = obj.empty();
 */