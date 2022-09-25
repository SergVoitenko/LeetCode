/*Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.



Example 1:

Input: s = "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()".
Example 2:

Input: s = ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()".
Example 3:

Input: s = ""
Output: 0


Constraints:

0 <= s.length <= 3 * 10 to the extent 4
s[i] is '(', or ')'.
*/

package main

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Проверяем пару скобок в конце строки",
			args: args{s: "(()"},
			want: 2,
		},
		{
			name: "Проверяем две пары самостоятельных закрывающихся скобок",
			args: args{s: ")()())"},
			want: 4,
		},
		{
			name: "Пустая строка",
			args: args{s: ""},
			want: 0,
		},
		{
			name: "Нет комбинации открывающихся и закрывающихся скобок",
			args: args{s: ")))))))))((((((((((("},
			want: 0,
		},
		{
			name: "Две возможные комбинации. Выбирается большая",
			args: args{s: "((())))))))))))(((())))((((((()))))))()()()()"},
			want: 30,
		},
		{
			name: "Неполное закрытие левой скобки. Вариант 1.",
			args: args{s: "((())))))))))))(((())))((((((()))))))()()()()((()"},
			want: 30,
		},
		{
			name: "Неполное закрытие левой скобки. Вариант 2.",
			args: args{s: "((())(()("},
			want: 4,
		}, 
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
