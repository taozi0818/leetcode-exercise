//Given a string s, find the length of the longest substring without repeating
//characters.
//
//
// Example 1:
//
//
//Input: s = "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
//
//
// Example 2:
//
//
//Input: s = "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
//
//
// Example 3:
//
//
//Input: s = "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Notice that the answer must be a substring, "pwke" is a subsequence and not a
//substring.
//
//
//
// Constraints:
//
//
// 0 <= s.length <= 5 * 10⁴
// s consists of English letters, digits, symbols and spaces.
//
// Related Topics 哈希表 字符串 滑动窗口 👍 7679 👎 0

package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdfjlski"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) (r int) {
	length := len(s)

	if length == 0 {
		return 0
	}

	m := map[uint8]int{
		s[0]: 0,
	}

	r = 0
	leftCursor := 0
	rightCursor := 1

	for rightCursor < length {
		if i, ok := m[s[rightCursor]]; ok {
			t := i + 1

			if l := rightCursor - leftCursor; l > r {
				r = l
			}

			for leftCursor < t {
				delete(m, s[leftCursor])
				leftCursor++
			}
		}

		m[s[rightCursor]] = rightCursor

		rightCursor++
	}

	end := rightCursor - leftCursor

	if end > r {
		return end
	}

	return r
}

//leetcode submit region end(Prohibit modification and deletion)
