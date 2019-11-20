package main

import (
	"errors"
	"fmt"
)

func main() {
	//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
	//输出：7 -> 0 -> 8
	//原因：342 + 465 = 807
	l1 := []int{2}
	l2 := []int{5, 6}
	head1, _ := createLink(l1)
	//printLink(head1)
	head2, _ := createLink(l2)
	sum := addTwoNumbers(head1, head2)
	printLink(sum)
	sum = reverseLink(sum)
	printLink(sum)
}

type ListNode struct {
    Val int
    Next *ListNode
}

func createLink(arr []int) (*ListNode, error) {
	head := new(ListNode)
	index := head
	if len(arr) == 0 {
		return head.Next, errors.New("list is null! ")
	}
	for _, value := range arr {
		tmp := &ListNode{Val:value}
		index.Next = tmp
		index = tmp
	}
	return head.Next, nil
}

func printLink(l *ListNode)  {
	tmp := l
	for ; tmp != nil ;  {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	index := head
	flag := 0
	// 在循环中判断元素为null进行赋值比在外面查找列表长度要简单
	for ; l1 != nil || l2 != nil; {
		i, j, sum := 0, 0, 0
		if l1 == nil {
			i = 0
		}else {
			i = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			j = 0
		}else {
			j = l2.Val
			l2 = l2.Next
		}
		// 判断进位时需要加上上一轮计算的进位
		if i + j + flag - 10 >= 0 {
			sum = i + j + flag - 10
			// 注意flag归位的位置
			flag = 1
		} else {
			sum = i + j + flag
			flag = 0
		}
		tmp := new(ListNode)
		tmp.Val = sum
		index.Next = tmp
		index = tmp
	}
	if flag > 0 {
		index.Next = &ListNode{Val:flag, Next:nil}
	}
	return head.Next
}

func reverseLink(l1 *ListNode) *ListNode {
	head := new(ListNode)
	for ; l1 != nil; l1 = l1.Next {
		tmp := &ListNode{Val:l1.Val, Next:nil}
		// 指针赋值，tmp.Next之后原链表就断了，所以需要new一个新的节点
		//tmp := l1
		tmp.Next = head.Next
		head.Next = tmp
	}
	return head.Next
}
