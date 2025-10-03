package main

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		data string
		want string
	}{
		{name: "standart-1", data: "abc", want: "cba"},
		{name: "phrase", data: "not yours", want: "sruoy ton"},
		{name: "phrase-with-utf8", data: "Hello, 世界", want: "界世 ,olleH"},
		{name: "spaces", data: "   jl   l", want: "l   lj   "},
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
		{name: "standart-2", data: "abb", want: false},
		{name: "diff case", data: "Madam", want: true},
		{name: "spaces&case", data: "Was it a car or a cat I saw", want: true},
		{name: "punctuation+spaces+case", data: "A man, a plan, a canal: Panama", want: true},
		{name: "emoji", data: "😂😊😂", want: true},
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

func TestSerialize(t *testing.T) {
	type obj struct {
		Num int `json:"num"`
	}

	js := JSONSerializer{}

	tests := []struct {
		name string
		data any
		want []byte
	}{
		{name: "standart-1", data: obj{Num: 234}, want: []byte(`{"num":234}`)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := js.Serialize(test.data)
			if err != nil {
				t.Fatal(err)
			}
			if string(test.want) != string(got) {
				t.Fatalf("want -> %s got -> %s", test.want, got)
			}
		})
	}
}
