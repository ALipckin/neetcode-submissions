/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    if head == nil || head.Next == nil {
        return
    }

    // 1. Находим середину
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    // 2. Разделяем список
    second := slow.Next
    slow.Next = nil

    // Reverse второй половины
    var prev *ListNode
    cur := second

    for cur != nil {
        next := cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }

    second = prev

    // 3. Слияние
    first := head

    for second != nil {
        tmp1 := first.Next
        tmp2 := second.Next

        first.Next = second
        second.Next = tmp1

        first = tmp1
        second = tmp2
    }
}