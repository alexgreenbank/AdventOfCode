// AoC 2017 day 1

// --- Day 1: Inverse Captcha ---
// 
// The night before Christmas, one of Santa's Elves calls you in a panic. "The printer's broken! We can't print the Naughty or Nice List!" By the time you make it to sub-basement 17, there are only a few minutes until midnight. "We have a big problem," she says; "there must be almost fifty bugs in this system, but nothing else can print The List. Stand in this square, quick! There's no time to explain; if you can convince them to pay you in stars, you'll be able to--" She pulls a lever and the world goes blurry.
// 
// When your eyes can focus again, everything seems a lot more pixelated than before. She must have sent you inside the computer! You check the system clock: 25 milliseconds until midnight. With that much time, you should be able to collect all fifty stars by December 25th.
// 
// Collect stars by solving puzzles. Two puzzles will be made available on each day millisecond in the advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!
// 
// You're standing in a room with "digitization quarantine" written in LEDs along one wall. The only door is locked, but it includes a small interface. "Restricted Area - Strictly No Digitized Users Allowed."
// 
// It goes on to explain that you may only leave by solving a captcha to prove you're not a human. Apparently, you only get one millisecond to solve the captcha: too fast for a normal human, but it feels like hours to you.
// 
// The captcha requires you to review a sequence of digits (your puzzle input) and find the sum of all digits that match the next digit in the list. The list is circular, so the digit after the last digit is the first digit in the list.
// 
// For example:
// 
//     1122 produces a sum of 3 (1 + 2) because the first digit (1) matches the second digit and the third digit (2) matches the fourth digit.
//     1111 produces 4 because each digit (all 1) matches the next.
//     1234 produces 0 because no digit matches the next.
//     91212129 produces 9 because the only digit that matches the next one is the last digit, 9.
// 
// What is the solution to your captcha?

package main

import (
	"strconv"
	"fmt"
	"bufio"
	"os"
)

func aoc2017day1a( text string ) int {
	// If input is shorter than 2 characters than it cannot have a non-zero score
	if len( text ) < 2 {
		return 0
	}
	tot := 0
	// Loop through string and check each character against the next
	for i := 0; i < len(text)-1; i++ {
		if text[i] == text[i+1] {
			n, _ := strconv.Atoi( string(text[i]) )
			tot += n
		}
	}
	// Now handle the first/last characters as the string is circular
	// Only if the string is longer than 2 characters otherwise we've already tested the first against the last
	if len(text) > 2 && text[0] == text[len(text)-1] {
		n, _ := strconv.Atoi( string(text[0]) )
		tot += n
	}
	return( tot )
}

func main() {
	testdata := map[string]int {
		"11": 1,		// simple match
		"12": 0,		// simple non-match
		"1122": 3,		// provided example
		"1111": 4,		// provided example
		"1234": 0,		// provided example
		"91212129": 9,		// provided example
	}
	for text, expected := range testdata {
		tot := aoc2017day1a( text )
		if tot != expected {
			fmt.Println( "ERROR: ["+text+"] returned "+strconv.Itoa(tot)+" and not expected value "+strconv.Itoa(expected) )
		}
	}
	// Read in the strings from stdin and process them
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		tot := aoc2017day1a( text )
		fmt.Println( tot )
	}
}
