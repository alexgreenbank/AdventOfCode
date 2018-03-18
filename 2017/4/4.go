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

package main

import (
	"strings"
	"os"
	"bufio"
	"fmt"
)

func main() {
        scanner := bufio.NewScanner(os.Stdin)
	words_ok := 0
	// Read in a line
        for scanner.Scan() {
                input := scanner.Text()
		seen := make( map[string]bool )
		ok := true
		// Parse the individual items within the line, space separated
		// Store in a map, reject if we see one that already exists
		for _, str := range strings.Split( input, " " ) {
			_, prs := seen[str]
			if prs  {
				// fmt.Println( "NO: ["+input+"] because of dup ["+str+"]" )
				ok=false
			}
			seen[str]=true
		}
		if ok {
			// fmt.Println( "YES: ["+input+"]" )
			words_ok++
		}
	}
	fmt.Println( words_ok )
}
