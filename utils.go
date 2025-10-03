package main

import "encoding/json"

type JSONSerializer struct{}

func (js JSONSerializer) Serialize(d any) ([]byte, error) {
	return json.Marshal(d)
}

func Reverse(s string) string {
	buf := []rune(s)
	st, end := 0, len(buf)-1

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

		if toLower(s[st]) != toLower(s[end]) {
			return false
		}

		st++
		end--
	}

	return true
}

type symbol interface {
	rune | byte
}

func isLetter[T symbol](ch T) bool {
	if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
		return true
	}

	return false
}

func toLower[T symbol](ch T) T {
	if isLetter(ch) && ch < 'a' {
		ch += 32
	}

	return ch
}
