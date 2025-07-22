package main

import (
	"strconv"
	"strings"
)

func singleNumber(nums []int) int {
	mapNum := make(map[int]int)
	answer := 0

	for _, num := range nums {
		if _, exists := mapNum[num]; exists {
			mapNum[num]++
		} else {
			mapNum[num] = 1
		}
	}

	for k, v := range mapNum {
		if v == 1 {
			answer = k
			break
		}
	}
	return answer
}

func isPalindrome(x int) bool {
	palindrome := strconv.Itoa(x)
	length := len(palindrome)

	for i := 0; i < length/2; i++ {
		if palindrome[i] != palindrome[length-i-1] {
			return false
		}
	}
	return true
}

func isValid(s string) bool {
	length := len(s)
	if length%2 != 0 {
		return false
	}
	pair := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}

	for i := 0; i < length; i++ {

		if i >= length/2 {
			break
		}
		aa := string(s[i])

		if _, exists := pair[aa]; !exists {
			return false
		}

		pairOne := pair[string(s[i])]
		last := string(s[length-i-1])
		next := string(s[i+1])

		if next == pairOne {
			i++
			continue
		}

		if last != pairOne {
			return false
		}
	}

	return true
}

func longestCommonPrefix(strs []string) string {
	shortest := ""

	for i := 0; i < len(strs); i++ {
		if i == 0 || len(strs[i]) < len(shortest) {
			shortest = strs[i]
		}
	}

	longest := shortest
	for i := 0; i < len(strs); i++ {
		for j := 0; j < len(shortest); j++ {
			if !strings.HasPrefix(strs[i], longest) {
				longest = longest[:len(longest)-1]
			}
		}
	}

	return longest
}

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}
	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}

func removeDuplicates(nums []int) int {
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func twoSum(nums []int, target int) []int {
	prevNums := map[int]int{}
	for i, num := range nums {
		targetNum := target - num
		targetNumIndex, ok := prevNums[targetNum]
		if ok {
			return []int{targetNumIndex, i}
		} else {
			prevNums[num] = i
		}
	}
	return []int{}
}

func main() {

}
