package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

func main() {
	one := make(chan string, 5)
	one <- "Hello world. This is First Test."
	one <- "hello World,  this Is. First  Test."
	two := make(chan string, 5)
	two <- "....WOW,,,,"
	two <- "Something Interesting Here"
	three := make(chan string, 5)
	three <- "String ....with Many   ,.  Spaces    "
	three <- "Simple String Without Lowercase Letters"
	close(one)
	close(two)
	close(three)
	out := mergeChannels(one, two, three)
	for value := range out {
		fmt.Println(getCapitalLetterWords(value))
	}
}

func getCapitalLetterWords(s string) string {
	var ans strings.Builder

	splitFunc := func(c rune) bool {
		return unicode.IsSpace(c) || c == ',' || c == '.'
	}
	splitted := strings.FieldsFunc(s, splitFunc)

	for idx, word := range splitted {
		if unicode.IsUpper(rune(word[0])) {
			ans.WriteString(word)
			if idx != len(splitted)-1 {
				ans.WriteString(" ")
			}
		}
	}
	return ans.String()
}

func mergeChannels(channels ...chan string) chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	for _, ch := range channels {
		wg.Add(1)
		go func(c chan string) {
			defer wg.Done()
			for value := range c {
				out <- value
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
