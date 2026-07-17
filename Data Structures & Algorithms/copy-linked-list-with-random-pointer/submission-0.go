/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }

    mp := make(map[*Node]*Node)

    for cur := head; cur != nil; cur = cur.Next {
        mp[cur] = &Node{
            Val: cur.Val,
        }
    }

    for cur := head; cur != nil; cur = cur.Next {
        copy := mp[cur]
        copy.Next = mp[cur.Next]
        copy.Random = mp[cur.Random]
    }

    return mp[head]
}
