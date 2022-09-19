/*
You are given a string s and an array of strings words. All the strings of words are of the same length.

A concatenated substring in s is a substring that contains all the strings of any permutation of words concatenated.

For example, if words = ["ab","cd","ef"], then "abcdef", "abefcd", "cdabef", "cdefab", "efabcd", and "efcdab" are all concatenated strings. "acdbef" is not a concatenated substring because it is not the concatenation of any permutation of words.
Return the starting indices of all the concatenated substrings in s. You can return the answer in any order.

Example 1:

Input: s = "barfoothefoobarman", words = ["foo","bar"]
Output: [0,9]
Explanation: Since words.length == 2 and words[i].length == 3, the concatenated substring has to be of length 6.
The substring starting at 0 is "barfoo". It is the concatenation of ["bar","foo"] which is a permutation of words.
The substring starting at 9 is "foobar". It is the concatenation of ["foo","bar"] which is a permutation of words.
The output order does not matter. Returning [9,0] is fine too.
Example 2:

Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
Output: []
Explanation: Since words.length == 4 and words[i].length == 4, the concatenated substring has to be of length 16.
There is no substring of length 16 is s that is equal to the concatenation of any permutation of words.
We return an empty array.
Example 3:

Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
Output: [6,9,12]
Explanation: Since words.length == 3 and words[i].length == 3, the concatenated substring has to be of length 9.
The substring starting at 6 is "foobarthe". It is the concatenation of ["foo","bar","the"] which is a permutation of words.
The substring starting at 9 is "barthefoo". It is the concatenation of ["bar","the","foo"] which is a permutation of words.
The substring starting at 12 is "thefoobar". It is the concatenation of ["the","foo","bar"] which is a permutation of words.

Constraints:

1 <= s.length <= 104
1 <= words.length <= 5000
1 <= words[i].length <= 30
s and words[i] consist of lowercase English letters.
*/
package main

import (
	"fmt"
	"sort"
	"strings"
	//	"os"
	//
	// "strings"
)

func main() {
	//s := "aaaaaaaaaaaaaa"
	//words := []string{"aa", "aa"}

	//Expected
	//[0,1,2,3,4,5,6,7,8,9,10]

	//s := "wordgoodgoodgoodbestword"
	// words := []string{"word", "good", "best", "word"}
	//Output: []

	//s := "barfoothefoobarman"
	//words := []string{"foo", "bar"}
	//Output: [0,9]

	//s := "lingmindraboofooowingdingbarrwingmonkeypoundcake"
	//words := []string{"fooo", "barr", "wing", "ding", "wing"}
	// Output: [13]

	s := "barfoofoobarthefoobarman"
	words := []string{"bar", "foo", "the"}
	//Output: [6,9,12]
	out := findSubstring(s, words)
	fmt.Println(out)

}
func findSubstring(s string, words []string) []int {
	type word struct {
		pos  int
		len  int
		word string
		find bool
	}
	comb := make([]word, len(words))
	var outPos []int
	lenWords := 0
	for i, w := range words {
		in := strings.Index(s, w)
		if in == -1 {
			return outPos
		} else {
			var w1 word
			w1.len = len(w)
			w1.pos = in
			w1.word = w
			comb[i] = w1
			lenWords += len(w)
		}
	}
	sort.Slice(comb, func(i, j int) bool {
		return comb[i].pos < comb[j].pos
	})
	delta := 0
	for len(s) >= lenWords {
		w := comb[0]
		comb[0].find = true
		strnew := s[w.pos+w.len:]
		for k := 1; k <= len(words)-1; k++ {
			for i := 1; i <= len(comb)-1; i++ {
				if !comb[i].find {
					if strings.HasPrefix(strnew, comb[i].word) {
						comb[i].find = true
						strnew = strnew[comb[i].len:]
						break
					}
				}
			}
		}
		wordOk := true
		for _, w1 := range comb {
			if !w1.find {
				wordOk = false
			}
		}
		if wordOk {
			outPos = append(outPos, w.pos+delta)
		}
		if w.pos == 0 {
			delta++
			s = s[1:]
		} else {
			delta += w.pos + 1
			s = s[w.pos+1:]
		}
		for i := range comb {
			in := strings.Index(s, comb[i].word)
			if in == -1 {
				comb[i].pos = in
				break
			} else {
				comb[i].find = false
				comb[i].pos = in
			}
		}
		sort.Slice(comb, func(i, j int) bool {
			return comb[i].pos < comb[j].pos
		})
		if comb[0].pos == -1 {
			break
		}
	}

	return outPos
}
