package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func validPassphrase(words []string) bool {
	var valid bool
	for i, _ := range words {
		for j, _ := range words {
			if i == j {
				continue
			}
			if words[i] == words[j] {
				return false
			}
		}
		valid = true
	}
	return valid
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func validAnagrams(words []string) bool {
	for i, _ := range words {
		for j, _ := range words {
			if i == j {
				continue
			}
			if SortString(words[i]) == SortString(words[j]) {
				return false
			}
		}
	}
	return true
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	var numValid, numAnag int

	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		strings := strings.Split(scanner.Text(), " ")
		fmt.Printf("%q\n", strings)
		if validPassphrase(strings) {
			numValid++
		}
		if validAnagrams(strings) {
			numAnag++
		}
	}

	fmt.Println("Valid passphrases:", numValid)
	fmt.Println("Valid anagrams:", numAnag)
}
