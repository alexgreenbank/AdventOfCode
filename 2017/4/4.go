// --- Day 4: High-Entropy Passphrases ---
// 
// A new system policy has been put in place that requires all accounts to use a passphrase instead of simply a password. A passphrase consists of a series of words (lowercase letters) separated by spaces.
// 
// To ensure security, a valid passphrase must contain no duplicate words.
// 
// For example:
// 
//     aa bb cc dd ee is valid.
//     aa bb cc dd aa is not valid - the word aa appears more than once.
//     aa bb cc dd aaa is valid - aa and aaa count as different words.
// 
// The system's full passphrase list is available as your puzzle input. How many passphrases are valid?

// --- Part Two ---
// 
// For added security, yet another system policy has been put in place. Now, a valid passphrase must contain no two words that are anagrams of each other - that is, a passphrase is invalid if any word's letters can be rearranged to form any other word in the passphrase.
// 
// For example:
// 
//     abcde fghij is a valid passphrase.
//     abcde xyz ecdab is not valid - the letters from the third word can be rearranged to form the first word.
//     a ab abc abd abf abj is a valid passphrase, because all letters need to be used when forming another word.
//     iiii oiii ooii oooi oooo is valid.
//     oiii ioii iioi iiio is not valid - any of these words can be rearranged to form any other word.
// 
// Under this new system policy, how many passphrases are valid?

package main

import (
	"strings"
	"sort"
	"os"
	"bufio"
	"fmt"
)

func SortString( s string ) string {
	// Sort a string by splitting it up, sorting the slice and then joining that result
	// Source: https://stackoverflow.com/questions/22688651/golang-how-to-sort-string-or-byte
	ss := strings.Split( s, "" )
	sort.Strings( ss )
	return strings.Join( ss, "" )
}

func main() {
        scanner := bufio.NewScanner(os.Stdin)
	// Keep track of the words valid under each set of rules
	words_ok_a := 0
	words_ok_b := 0
	// Read in a line
        for scanner.Scan() {
                input := scanner.Text()
		seen_a := make( map[string]bool )
		seen_b := make( map[string]bool )
		ok_a := true
		ok_b := true
		// Parse the individual items within the line, space separated
		for _, str := range strings.Split( input, " " ) {
			// Store in a map, reject if we see one that already exists
			_, prs := seen_a[str]
			if prs {
				ok_a=false
			}
			seen_a[str]=true
			// For part b we sort the letters in each substring first
			// That way we automatically search for any anagrams
			sstr := SortString( str )
			_, prs = seen_b[sstr]
			if prs {
				ok_b=false
			}
			seen_b[sstr]=true
		}
		if ok_a {
			words_ok_a++
		}
		if ok_b {
			words_ok_b++
		}
	}
	fmt.Println( words_ok_a )
	fmt.Println( words_ok_b )
}
