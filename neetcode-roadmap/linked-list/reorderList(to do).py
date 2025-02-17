# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        first, last = head 
        counter = 0
        while last.next != None:
            counter += 1
            last = last.next 

        half = head
        i = 0
        while i < counter / 2:
            i += 1
            half = half.next

        prev = None
        curr = half 
        while curr:
            tmp = curr.next
            curr.next = prev
            prev = curr 
            curr = tmp

        i = 0
        p1, p2 = head, half 
        curr = head 
        while p1 and p2:
            tmp = curr.next
            if i % 2 == 0:
                curr.next = p2 


