import	"unicode"


func isPalindrome(s string) bool {
	r := []rune(s)
	sl := []rune{}
	for i := 0; i < len(r); i++ {
		if unicode.IsDigit(r[i]) || unicode.IsLetter(r[i]) {
			sl = append(sl, unicode.ToLower(r[i]))
		}
	}
	min := 0
	max := len(sl) - 1
	for min < max {
		if sl[min] == sl[max] {
			min++
			max--
		} else {
			return false
		}
	}
	return true

}