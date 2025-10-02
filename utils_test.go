package main

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		data string
		want string
	}{
		{name: "standart-1", data: "abc", want: "cba"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Reverse(test.data)
			if test.want != got {
				t.Fatalf("want -> %s got -> %s", test.want, got)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		data string
		want bool
	}{
		{name: "standart-1", data: "aba", want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsPalindrome(test.data)
			if test.want != got {
				t.Fatalf("want -> %v got -> %v", test.want, got)
			}
		})
	}
}
