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

// --- Part Two ---
// 
// "Great work; looks like we're on the right track after all. Here's a star for your effort." However, the program seems a little worried. Can programs be worried?
// 
// "Based on what we're seeing, it looks like all the User wanted is some information about the evenly divisible values in the spreadsheet. Unfortunately, none of us are equipped for that kind of calculation - most of us specialize in bitwise operations."
// 
// It sounds like the goal is to find the only two numbers in each row where one evenly divides the other - that is, where the result of the division operation is a whole number. They would like you to find those numbers on each line, divide them, and add up each line's result.
// 
// For example, given the following spreadsheet:
// 
// 5 9 2 8
// 9 4 7 3
// 3 8 6 5
// 
//     In the first row, the only two numbers that evenly divide are 8 and 2; the result of this division is 4.
//     In the second row, the two numbers are 9 and 3; the result is 3.
//     In the third row, the result is 2.
// 
// In this example, the sum of the results would be 4 + 3 + 2 = 9.
// 
// What is the sum of each row's result in your puzzle input?

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
	tota := 0
	totb := 0
	for scanner.Scan() {
		text := scanner.Text()
		var list = []int{}
		// Each line has a list of tab separated non-negative integers
		// Grab them and stick them in the list slice
		for _, str := range strings.Split( text, "\t" ) {
			// fmt.Println( str )
			nos, _ := strconv.Atoi( str )
			if nos < 0 {
				fmt.Println( "ERROR: unexpected negative number ["+str+"]" )
			}
			list = append( list, nos )
		}
		// Part a
		// 'Score' for each line is max-min values
		min := 0
		max := 0
		for _, nos := range list {
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
		tota += max-min
		// Part b
		// 'Score' for each line is the result of only two numbers that divide exactly
		// O(n^2) this part - no other way to do it
		partb := 0
		for _, nosa := range list {
			for _, nosb := range list {
				if nosa > nosb && ( nosa % nosb == 0 ) {
					if partb != 0 {
						fmt.Println( "2017_2_b: More than one divisor on line ["+text+"]" )
					}
					partb = nosa / nosb
				}
			}
		}
		if partb == 0 {
			fmt.Println( "2017_2_b: No divisor found in line ["+text+"]" )
		}
		totb += partb
		
	}
	fmt.Println( tota )
	fmt.Println( totb )
}
