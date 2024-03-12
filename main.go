package main

func main() {
	input := "babad"
	println("Example 1:", input)
	LongestPalindrome(input)
	input = "cbbd"
	println("Example 2:", input)
	LongestPalindrome(input)
}

// IsPalindrome function have runtime of O(n)
func IsPalindrome(input string) bool {
	if len(input) == 1 {
		return true
	}
	lastIndex := len(input) - 1
	for i := 0; i <= lastIndex/2 && i < (lastIndex-i); i++ {
		if input[i] != input[lastIndex-i] {
			return false
		}
	}
	return true
}

// LongestPalindrome will print all unique string that are the longest Palindrome
// and have a runtime of O(n^3)
func LongestPalindrome(input string) {
	// If the input string are 1 length it a Palindrome by fact
	if len(input) == 1 {
		println(input)
	}
	list := make([]string, 0)
	longest := 0
	for i := 0; i < len(input); i++ {
		for j := i + 1; j <= len(input); j++ {
			str := input[i:j]
			if IsPalindrome(str) && len(str) > longest {
				longest = len(str)
				list = nil
				list = append(list, str)
			} else if IsPalindrome(str) && len(str) == longest {
				list = append(list, str)
			}
		}
	}
	if len(list) > 1 {
		for _, strInList := range list {
			println(strInList)
		}
		return
	}
	println(list[0])
}
