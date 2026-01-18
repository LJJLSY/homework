package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 只出现一次的数字
func ControlFlow() {
	fmt.Println("--------只出现一次的数字--------")
	arr := [...]int{2, 2, 3, 3, 4, 4, 4, 5}
	m := make(map[int]int)

	for _, v := range arr {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	fmt.Println(m)

	for k, v := range m {
		if v == 1 {
			fmt.Printf("数组中只出现一次的元素是：%d\n", k)
		}
	}
}

// 回文数
func Palindrome(x int) {
	fmt.Print("\n--------回文数--------\n")
	if x < 0 {
		fmt.Printf("数字 %d 不是回文数\n", x)
		return
	}

	s := strconv.Itoa(x)
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			fmt.Printf("数字 %d 不是回文数\n", x)
			return
		}
		left++
		right--
	}

	fmt.Printf("数字 %d 是回文数\n", x)
}

// 有效的括号
func ValidBrackets(s string) {
	fmt.Print("\n--------有效的括号--------\n")
	arr := []rune{}
	m := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			arr = append(arr, char)
		case ')', '}', ']':
			if len(arr) == 0 {
				fmt.Printf("字符串 %s 不含有效括号\n", s)
				return
			}
			top := arr[len(arr)-1]
			arr = arr[:len(arr)-1]

			if top != m[char] {
				fmt.Printf("字符串 %s 不含有效括号\n", s)
				return
			}
		}
	}

	if len(arr) == 0 {
		fmt.Printf("字符串 %s 含有效括号\n", s)
	} else {
		fmt.Printf("字符串 %s 不含有效括号\n", s)
	}
}

// 最长公共前缀
func CommonPrefix() {
	fmt.Println("--------最长公共前缀--------")
	s := [...]string{"flower", "flow", "flight"}

	for i, c1 := range s[0] {
		for _, v := range s[1:] {
			if len(v) > i {
				c2 := v[i]
				if string(c1) != string(c2) {
					result := s[0][:i]
					fmt.Print(result)
					return
				}
			} else {
				result := s[0][:i]
				fmt.Print(result)
				return
			}
		}
	}
	result := s[0]
	fmt.Print(result)
}

// 加一
func increment() {
	fmt.Println("--------加一--------")
	arr := []int{9, 9, 9, 9}

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != 9 {
			arr[i] += 1
			strarr := strings.ReplaceAll(fmt.Sprintf("%v", arr), " ", ",")
			fmt.Printf("加一后的数组%s\n", strarr)
			return
		} else {
			arr[i] = 0
		}
	}

	arr2 := make([]int, len(arr)+1)
	arr2[0] = 1
	strarr2 := strings.ReplaceAll(fmt.Sprintf("%v", arr2), " ", ",")
	fmt.Printf("加一后的数组%s\n", strarr2)
}

// 删除有序数组中的重复项
func Duplicates() {
	fmt.Println("--------删除有序数组中的重复项--------")
	nums := [...]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	i := 0

	for _, v := range nums[1:] {
		if v != nums[i] {
			i++
			nums[i] = v
		}
	}
	fmt.Printf("新长度：%d\n", i+1)
}

// 合并区间
func MergeIntervals() {
	fmt.Println("--------合并区间--------")
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}
	i := 0
	for _, v := range intervals[1:] {
		if v[0] <= merged[i][1] && v[1] > merged[i][1] {
			merged[i][1] = v[1]
		} else if v[0] > merged[i][1] {
			merged = append(merged, v)
			i++
		}
	}
	fmt.Printf("合并后的区间数组：%v\n", merged)
}

// 两数之和
func NumSum() {
	fmt.Println("--------两数之和--------")
	nums := [...]int{3, 3}
	target := 6
	m := make(map[int]int)

	for i, v := range nums {
		complement := target - v
		if j, ok := m[complement]; ok {
			m1 := strings.ReplaceAll(fmt.Sprintf("%v", []int{j, i}), " ", ",")
			fmt.Printf("两数索引：%s\n", m1)
			return
		}
		m[v] = i
	}

	return
}

func main() {
	ControlFlow()
	Palindrome(1221)
	ValidBrackets("(()[]){}")
	CommonPrefix()
	increment()
	Duplicates()
	MergeIntervals()
	NumSum()
}
