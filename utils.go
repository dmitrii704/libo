package main

func Reverse(s string) string {
	st, end := 0, len(s)-1
	buf := []byte(s)

	for st < end {
		buf[st], buf[end] = buf[end], buf[st]

		st++
		end--
	}

	return string(buf)
}

func IsPalindrome(s string) bool {
	st, end := 0, len(s)-1

	for st < end {
		if !isLetter(s[st]) {
			st++
			continue
		}

		if !isLetter(s[end]) {
			end--
			continue
		}

		if s[st] != s[end] {
			return false
		}

		st++
		end--
	}

	return true
}

func isLetter(ch byte) bool {
	if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
		return true
	}

	return false
}
