package main

/**
给你一个头结点为 head 的单链表和一个整数 k ，请你设计一个算法将链表分隔为 k 个连续的部分。
每部分的长度应该尽可能的相等：任意两部分的长度差距不能超过 1 。这可能会导致有些部分为 null 。
这 k 个部分应该按照在链表中出现的顺序排列，并且排在前面的部分的长度应该大于或等于排在后面的长度。
返回一个由上述 k 部分组成的数组。

示例 1：
输入：head = [1,2,3], k = 5
输出：[[1],[2],[3],[],[]]
解释：
第一个元素 output[0] 为 output[0].val = 1 ，output[0].next = null 。
最后一个元素 output[4] 为 null ，但它作为 ListNode 的字符串表示是 [] 。

示例 2：
输入：head = [1,2,3,4,5,6,7,8,9,10], k = 3
输出：[[1,2,3,4],[5,6,7],[8,9,10]]
解释：
输入被分成了几个连续的部分，并且每部分的长度相差不超过 1 。前面部分的长度大于等于后面部分的长度。

提示：
链表中节点的数目在范围 [0, 1000]
0 <= Node.val <= 1000
1 <= k <= 50
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	res := make([]*ListNode, k)
	cur := head
	n := 0
	for cur != nil {
		n++
		cur = cur.Next
	}

	mod, size := n%k, n/k

	cur = head
	for i := 0; i < k && cur != nil; i++ {
		res[i] = cur
		curSize := size
		if i < mod {
			curSize++
		}

		for i := 0; i < curSize-1; i++ {
			cur = cur.Next
		}

		cur, cur.Next = cur.Next, nil
	}
	return res
}

func main() {
	var head *ListNode = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{3, nil}}}
	splitListToParts(head, 5)
}

/** java

class ListNode {
    int val;
    ListNode next;

    ListNode() {
    }

    ListNode(int val) {
        this.val = val;
    }

    ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }
}

class Solution {
    public ListNode[] splitListToParts(ListNode head, int k) {
        ListNode[] res = new ListNode[k];

        //计算长度
        int n = 0;
        ListNode cur = head;
        while (cur != null) {
            n++;
            cur = cur.next;
        }

        //大小
        int size = n / k;
        //余数
        int mod = n % k;
        //cur指针指向头结点
        cur = head;
        for (int i = 0; i < k && cur != null; i++) {
            //每一段的头节点
            res[i] = cur;
            //将多余的节点平摊到前几段上
            int cursize = size + (mod-- > 0 ? 1 : 0);
            //分段
            for (int j = 0; j < cursize - 1; j++) {
                cur = cur.next;
            }
            //截断
            ListNode temp = cur.next;
            cur.next = null;
            cur = temp;
        }
        return res;
    }
}

*/
