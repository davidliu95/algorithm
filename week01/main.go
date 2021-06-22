package main

import "fmt"

func main() {
	a := []int{-1, -1, 1}
	fmt.Println(subarraySum(a, 0))
}

//加一
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}
	var result []int
	for i := 0; i < len(digits)+1; i++ {
		if i == 0 {
			result = append(result, 1)
		} else {
			result = append(result, digits[i-1])
		}
	}
	return result
}

//合并两个有序链表
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var listNode *ListNode
	if l1.Val < l2.Val {
		listNode = &ListNode{
			Val:  l1.Val,
			Next: mergeTwoLists(l1.Next, l2),
		}
	} else {
		listNode = &ListNode{
			Val:  l2.Val,
			Next: mergeTwoLists(l1, l2.Next),
		}
	}
	return listNode
}

//环形链表
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

//和为K的子数组
func subarraySum(nums []int, k int) int {
	//前缀和               sum[i-1]+nums[i]
	//统计第l到第r个数的和   nums[r]-nums[l-1]
	sum := []int{0}
	for i := 0; i < len(nums); i++ {
		sum = append(sum, nums[i]+sum[i])
	}

	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if sum[j+1]-sum[i] == k {
				count++
			}
		}
	}
	return count
}

//两数之和
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := m[target-nums[i]]
		if ok {
			return []int{m[target-nums[i]], i}
		}
		m[nums[i]] = i
	}
	return []int{}
}

//盛最多水的容器
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	leftIndex := 0
	rightIndex := len(height) - 1
	area := 0
	for leftIndex < rightIndex {
		leftSide := height[leftIndex]
		rightSide := height[rightIndex]
		if leftSide < rightSide {
			if leftSide*(rightIndex-leftIndex) > area {
				area = leftSide * (rightIndex - leftIndex)
			}
			leftIndex += 1
		} else {
			if rightSide*(rightIndex-leftIndex) > area {
				area = rightSide * (rightIndex - leftIndex)
			}
			rightIndex -= 1
		}
	}
	return area
}
