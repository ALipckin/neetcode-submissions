func isPalindrome(s string) bool {
	s = normalize(s)
	li := len(s) - 1
	f, l := 0, li
	for f < l {
		if s[f] != s[l] {
			return false
		}
		f++
		l--
	}
	return true
}

func normalize(s string) string {
	var b []rune
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b = append(b, r)
		}
	}
	return string(b)
}
