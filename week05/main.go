package main

import "fmt"

func main() {
	//s:=[]int{1,2,3,4,5,6,7,8,9,10}
	//fmt.Println(shipWithinDays(s,5))
	//fmt.Println(minEatingSpeed([]int{30,11,23,4,20},6))
	fmt.Println(minEatingSpeed([]int{312884470}, 968709470))
	fmt.Println(30 / 17)
	fmt.Println(30 % 17)
}

//在 D 天内送达包裹的能力
func shipWithinDays(weights []int, days int) int {
	//首先要找出这堆包裹里的重量最大值，以及计算所有包裹总重量
	maxWeight, totalWeight := 0, 0
	for i := 0; i < len(weights); i++ {
		if maxWeight < weights[i] {
			maxWeight = weights[i]
		}
		totalWeight += weights[i]
	}
	//定义一个方法,用来计算每天运送x重量，需要多少天
	var calc func(x int) int
	calc = func(x int) int {
		currentWeight := 0
		result := 1
		for i := 0; i < len(weights); i++ {
			weight := weights[i]
			//如果当天重量大于等于需要运送的重量，就将当天重量清零，并将天数增加一天，否则就继续累计当天重量
			if currentWeight+weight > x {
				result++
				currentWeight = 0
			}
			currentWeight += weight
		}
		return result
	}
	//通过二分查找，找出满足要求的最低重量,左边界不能小于最大重量，右边界不能大于所有包裹重量总和
	left := maxWeight
	right := totalWeight
	for left < right {
		mid := (left + right) / 2

		if calc(mid) <= days {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

//爱吃香蕉的珂珂
func minEatingSpeed(piles []int, h int) int {
	//找出数组中香蕉堆最大的值
	max := 0
	for i := 0; i < len(piles); i++ {
		if max < piles[i] {
			max = piles[i]
		}
	}
	//如果时间刚好等于数组长度，返回最大元素
	if len(piles) == h {
		return max
	}
	//定义一个函数，传入吃香蕉的速度值x,计算吃完所有香蕉，需要多少小时
	var calc func(x int) int
	calc = func(x int) int {
		//最终消耗时间
		result := 0
		if x > 0 {
			for i := 0; i < len(piles); i++ {
				//计算第i堆香蕉，按x的速度消耗了多少小时
				result += piles[i] / x
				//如果有余数，小时加1
				if piles[i]%x > 0 {
					result++
				}
			}
		}
		return result
	}
	//通过二分查找，找出满足要求的最小速度，左边界为1，右边界不超过数组中最大值
	left := 1
	right := max
	for left <= right {
		mid := (left + right) / 2
		if calc(mid) <= h {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
