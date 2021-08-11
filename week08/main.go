package main

import (
	"fmt"
	"strings"
)

func main() {
	//s:=[]string{"flower","flow","flight"}
	//longestCommonPrefix(s)

	fmt.Println(reverseWords(" asdasd df f"))
}

//最长公共前缀
func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		//拿到第一个字符串，循环后面的每个字符串，找出公共前缀
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

//最长公共前缀（二分法）
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	isCommonPrefix := func(length int) bool {
		str0, count := strs[0][:length], len(strs)
		for i := 1; i < count; i++ {
			if strs[i][:length] != str0 {
				return false
			}
		}
		return true
	}
	minLength := len(strs[0])
	for _, s := range strs {
		if len(s) < minLength {
			minLength = len(s)
		}
	}
	low, high := 0, minLength
	for low < high {
		mid := (high-low+1)/2 + low
		if isCommonPrefix(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return strs[0][:low]
}

//翻转字符串里的单词
func reverseWords(s string) string {

	//当只有一个单词时直接返回
	if !strings.Contains(s, " ") {
		return s
	}

	//创建一个存放单词的数组
	var wordList []string
	//wordBeginIndex 用于记录每个单词的第一个字母的下标索引
	wordBeginIndex := -1
	//循环整个字符串
	for i := 0; i < len(s); i++ {
		u := s[i : i+1]

		if u != " " {
			//不是空格时，且wordBeginIndex还是初始值，就记录当前索引为单词首字母下标索引
			if wordBeginIndex < 0 {
				wordBeginIndex = i
			}
			//如果索引来到最后一位，且wordBeginIndex存在单词首字母下标索引，就截取最后一个单词
			if i == len(s)-1 && wordBeginIndex > 0 {
				word := s[wordBeginIndex:]
				wordList = append(wordList, word)
			}
		} else {
			//如果当前字符是空格，且wordBeginIndex存在单词首字母下标索引，表明wordBeginIndex，到当前下标之前的字符是一个完整的单词，并存到单词数组中
			if wordBeginIndex >= 0 {
				word := s[wordBeginIndex:i]
				wordList = append(wordList, word)

				wordBeginIndex = -1
			}
		}
	}

	//最终倒叙拼接单词
	result := ""
	for i := 0; i < len(wordList); i++ {
		result = " " + wordList[i] + result
	}

	return result[1:]
}
