package easy

import "github.com/taydy/go-leetcode/util"

// 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
//		push(x) -- 将元素 x 推入栈中。
//		pop() -- 删除栈顶的元素。
//		top() -- 获取栈顶元素。
//		getMin() -- 检索栈中的最小元素。
//
//	示例:
//		MinStack minStack = new MinStack();
//		minStack.push(-2);
//		minStack.push(0);
//		minStack.push(-3);
//		minStack.getMin();   --> 返回 -3.
//		minStack.pop();
//		minStack.top();      --> 返回 0.
//		minStack.getMin();   --> 返回 -2.
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/min-stack
//
type Node struct {
	Value int
	Min   int
	Next  *Node
}

type MinStack struct {
	Head *Node
}

/** initialize your data structure here. */
func MinStackConstructor() MinStack {
	return MinStack{Head: nil}
}

func (m *MinStack) Push(x int) {
	if m.Head == nil {
		m.Head = &Node{
			Value: x,
			Min:   x,
		}
	} else {
		node := &Node{
			Value: x,
			Min:   util.MinInt(m.Head.Min, x),
			Next:  m.Head,
		}
		m.Head = node
	}
}

func (m *MinStack) Pop() {
	if m.Head != nil {
		m.Head = m.Head.Next
	}
}

func (m *MinStack) Top() int {
	if m.Head != nil {
		return m.Head.Value
	}
	return -1
}

func (m *MinStack) GetMin() int {
	return m.Head.Min
}
