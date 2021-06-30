package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := []string{"9001 discuss.leetcode.com"}
	fmt.Println(subdomainVisits(a))

	//fmt.Println("discuss.leetcode.com"[strings.Index("discuss.leetcode.com", ".")+1:])
}

//数组的度
func findShortestSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//创建一个以nums的元素为key的map
	resultMap := make(map[int]int)
	firstMap := make(map[int]int)
	lastMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		//第一个数直接计入resultMap
		if i == 0 {
			resultMap[nums[i]] = 1
		} else {
			//找到map中对应的key
			value, ok := resultMap[nums[i]]
			if ok {
				resultMap[nums[i]] = value + 1
			} else {
				resultMap[nums[i]] = 1
			}
		}
		//记下每个数的最小下标和最大下标
		_, ok := firstMap[nums[i]]
		if !ok {
			firstMap[nums[i]] = i
		}
		lastMap[nums[i]] = i
	}
	//找出计数最大值
	maxCount := 0
	for key := range resultMap {
		if resultMap[key] > maxCount {
			maxCount = resultMap[key]
		}
	}
	min := int(^uint(0) >> 1)
	for key := range resultMap {
		//找出统计数量和最大计数相同的值
		if resultMap[key] == maxCount {
			//找到当前值的最小下标和最大小标
			first := firstMap[key]
			last := lastMap[key]
			//计算它们之间的差,找到最小值
			if min != 0 && (last-first+1) < min {
				min = last - first + 1
			}
		}
	}
	return min
}

//子域名访问计数
func subdomainVisits(cpdomains []string) []string {
	if len(cpdomains) == 0 {
		return nil
	}
	//声明一个统计所有域名访问次数的map
	countMap := make(map[string]int)
	for i := 0; i < len(cpdomains); i++ {
		split := strings.Split(cpdomains[i], " ")
		count, _ := strconv.Atoi(split[0])
		domain := split[1]

		//根据"."分割域名
		domains := strings.Split(domain, ".")
		temp := ""
		//从根域名开始依次往上拼接，并将域名的次数统计到map中
		for i := len(domains) - 1; i >= 0; i-- {
			if i == len(domains)-1 {
				temp = domains[i]
			} else {
				temp = domains[i] + "." + temp
			}
			result, ok := countMap[temp]
			if ok {
				countMap[temp] = result + count
			} else {
				countMap[temp] = count
			}
		}
	}
	//最终将map中的所有域名放在数组中
	var result []string
	for key := range countMap {
		s := strconv.Itoa(countMap[key]) + " " + key
		result = append(result, s)
	}
	return result
}
