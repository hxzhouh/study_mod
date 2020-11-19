package string

import "strings"

var bc [256]int

type stringFinder struct {
	pattern        string
	badCharSkip    [256]int // 坏词典 hash值
	goodSuffixSkip []int    // 好的匹配串
}

func makeStringFinder(pattern string) *stringFinder {
	f := &stringFinder{
		pattern:        pattern,
		goodSuffixSkip: make([]int, len(pattern)),
	}
	// last is the index of the last character in the pattern.
	last := len(pattern) - 1

	// Build bad character table.
	// Bytes not in the pattern can skip one pattern's length.
	for i := range f.badCharSkip {
		f.badCharSkip[i] = len(pattern)
	}
	// The loop condition is < instead of <= so that the last byte does not
	// have a zero distance to itself. Finding this byte out of place implies
	// that it is not in the last position.
	for i := 0; i < last; i++ {
		f.badCharSkip[pattern[i]] = last - i
	}

	// Build good suffix table.
	// First pass: set each value to the next index which starts a prefix of
	// pattern.
	lastPrefix := last
	for i := last; i >= 0; i-- {
		if strings.HasPrefix(pattern, pattern[i+1:]) {
			lastPrefix = i + 1
		}
		// lastPrefix is the shift, and (last-i) is len(suffix).
		f.goodSuffixSkip[i] = lastPrefix + last - i
	}
	// Second pass: find repeats of pattern's suffix starting from the front.
	for i := 0; i < last; i++ {
		lenSuffix := longestCommonSuffix(pattern, pattern[1:i+1])
		if pattern[i-lenSuffix] != pattern[last-lenSuffix] {
			// (last-i) is the shift, and lenSuffix is len(suffix).
			f.goodSuffixSkip[last-lenSuffix] = lenSuffix + last - i
		}
	}

	return f
}
func longestCommonSuffix(a, b string) (i int) {
	for ; i < len(a) && i < len(b); i++ {
		if a[len(a)-1-i] != b[len(b)-1-i] {
			break
		}
	}
	return
}

func (f *stringFinder) next(text string) int {
	i := len(f.pattern) - 1
	for i < len(text) {
		// Compare backwards from the end until the first unmatching character.
		j := len(f.pattern) - 1
		for j >= 0 && text[i] == f.pattern[j] {
			i--
			j--
		}
		if j < 0 {
			return i + 1 // match
		}
		i += max(f.badCharSkip[text[i]], f.goodSuffixSkip[j])
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//func generateBC(b []rune, lengthB int) {
//	for i := 0; i < 256; i++ {
//		bc[i] = -1
//	}
//	for i := 0; i < lengthB; i++ {
//		bc[b[i]] = i // 记录每个字母最后一次出现的位置,
//	}
//	//fmt.Println(bc)
//}
//func bm(a, b []rune, lengthA, lengthB int) int {
//	generateBC(b, lengthB) // 初始化模式串 hash
//
//	i := 0                 // i表示主串与模式串对齐的第一个字符
//	for i <= lengthA-lengthB {
//		j := 0
//		for j = lengthB - 1; j >= 0; j -= 1 { // 从前开始匹配
//			if a[i+j] != b[j] { // 坏字符对应模式串中的下标是j
//				break
//			}
//		}
//		if j < 0 { //完事，匹配成功
//			return i
//
//		}
//		//向后滑动
//		i = i + (j - bc[a[i+j]])
//	}
//	return -1
//}
//
//func generateGS(b []rune, lengthB int, suffix, prefix *[]int) {
//	for i := 0; i < lengthB; i++ {
//		*suffix[i]=-1
//	}
//}
