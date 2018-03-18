// AoC 2017 day 2

// --- Day 2: Corruption Checksum ---
// 
// As you walk through the door, a glowing humanoid shape yells in your direction. "You there! Your state appears to be idle. Come help us repair the corruption in this spreadsheet - if we take another millisecond, we'll have to display an hourglass cursor!"
// 
// The spreadsheet consists of rows of apparently-random numbers. To make sure the recovery process is on the right track, they need you to calculate the spreadsheet's checksum. For each row, determine the difference between the largest value and the smallest value; the checksum is the sum of all of these differences.
// 
// For example, given the following spreadsheet:
// 
// 5 1 9 5
// 7 5 3
// 2 4 6 8
// 
//     The first row's largest and smallest values are 9 and 1, and their difference is 8.
//     The second row's largest and smallest values are 7 and 3, and their difference is 4.
//     The third row's difference is 6.
// 
// In this example, the spreadsheet's checksum would be 8 + 4 + 6 = 18.
// 
// What is the checksum for the spreadsheet in your puzzle input?

package main

import (
	"strconv"
	"strings"
	"fmt"
	"bufio"
	"os"
)

func main() {
	// Read in the strings from stdin and process them
	scanner := bufio.NewScanner(os.Stdin)
	// Can process the input one line at a time keeping track of the total checksum along the way
	tot := 0
	for scanner.Scan() {
		text := scanner.Text()
		min := 0
		max := 0
		// Each line has a list of tab separated non-negative integers
		for _, str := range strings.Split( text, "\t" ) {
			// fmt.Println( str )
			nos, _ := strconv.Atoi( str )
			if nos < 0 {
				fmt.Println( "ERROR: unexpected negative number ["+str+"]" )
			}
			// These will be both zero if we haven't initialised them yet
			if min == 0 && max == 0 {
				min=nos;
				max=nos;
			}
			// Otherwise check for new min/max
			if nos < min {
				min = nos
			}
			if nos > max {
				max = nos
			}
		}
		// Add line checksum value to total
		tot += max-min
	}
	fmt.Println( tot )
}
