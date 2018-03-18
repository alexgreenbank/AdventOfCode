// --- Day 3: Spiral Memory ---
// 
// You come across an experimental new kind of memory stored on an infinite two-dimensional grid.
// 
// Each square on the grid is allocated in a spiral pattern starting at a location marked 1 and then counting up while spiraling outward. For example, the first few squares are allocated like this:
// 
// 17  16  15  14  13
// 18   5   4   3  12
// 19   6   1   2  11
// 20   7   8   9  10
// 21  22  23---> ...
// 
// While this is very space-efficient (no squares are skipped), requested data must be carried back to square 1 (the location of the only access port for this memory system) by programs that can only move up, down, left, or right. They always take the shortest path: the Manhattan Distance between the location of the data and square 1.
// 
// For example:
// 
//     Data from square 1 is carried 0 steps, since it's at the access port.
//     Data from square 12 is carried 3 steps, such as: down, left, left.
//     Data from square 23 is carried only 2 steps: up twice.
//     Data from square 1024 must be carried 31 steps.
// 
// How many steps are required to carry the data from the square identified in your puzzle input all the way to the access port?

// --- Part Two ---
// 
// As a stress test on the system, the programs here clear the grid and then store the value 1 in square 1. Then, in the same allocation order as shown above, they store the sum of the values in all adjacent squares, including diagonals.
// 
// So, the first few squares' values are chosen as follows:
// 
//     Square 1 starts with the value 1.
//     Square 2 has only one adjacent filled square (with value 1), so it also stores 1.
//     Square 3 has both of the above squares as neighbors and stores the sum of their values, 2.
//     Square 4 has all three of the aforementioned squares as neighbors and stores the sum of their values, 4.
//     Square 5 only has the first and fourth squares as neighbors, so it gets the value 5.
// 
// Once a square is written, its value does not change. Therefore, the first few squares would receive the following values:
// 
// 147  142  133  122   59
// 304    5    4    2   57
// 330   10    1    1   54
// 351   11   23   25   26
// 362  747  806--->   ...
// 
// What is the first value written that is larger than your puzzle input?

package main

import (
	"strconv"
	"fmt"
	"os"
)

// Simple integer absolute function
func Iabs( val int ) int {
	if val < 0  {
		return( -val )
	}
	return( val )
}

func main() {
	if len( os.Args ) < 2 {
		fmt.Println( os.Args[0]+": <nos>" )
		return	
	}
	nos, _ := strconv.Atoi( os.Args[1] )

	// Store the nodes in a map linked by x,y co-ords
	type Key struct {
		x, y int
	}
	nodes := make(map[Key]int)

	// Initialise values
	val := 1
	nodes[Key{ 0, 0 }] = val
	x := 0
	y := 0
	dir := "E"

	for val < nos {
		// Move one more place on in current direction and increment value
		val++
		if dir == "E" {
			x+=1
		} else if dir == "N" {
			y+=1
		} else if dir == "W" {
			x-=1
		} else if dir == "S" {
			y-=1
		} else {
			fmt.Println( "ERROR, Unhandled direction ["+dir+"]" )
			return
		}
		// Set value at this location
		nodes[Key{ x, y }] = val
		// If spot to the 'left' of here is empty then we need to turn left
		if dir == "E" { // Check if value to 'left' (N) exists
			_, prs := nodes[Key{ x, y+1 }]
			if !prs { // No value, Turn left (N)
				dir = "N"
			}
		} else if dir == "N" { // Check if value to 'left' (W) exists
			_, prs := nodes[Key{ x-1, y }]
			if !prs { // No value, Turn left (W)
				dir = "W"
			}
		} else if dir == "W" { // Check if value to 'left' (S) exists
			_, prs := nodes[Key{ x, y-1 }]
			if !prs { // No value, Turn left (S)
				dir = "S"
			}
		} else if dir == "S" { // Check if value to 'left' (E) exists
			_, prs := nodes[Key{ x+1, y }]
			if !prs { // No value, Turn left (E)
				dir = "E"
			}
		}
	}
	// Output the answer for part a
	fmt.Println( Iabs(x)+Iabs(y) )

	// Clean out nodes map and reset x,y,dir
	nodes = make(map[Key]int)
	nodes[Key{ 0, 0 }] = 1
	x = 0
	y = 0
	dir = "E"

	val = 1
	for val < nos {
		// Move one more place on in current direction
		if dir == "E" {
			x+=1
		} else if dir == "N" {
			y+=1
		} else if dir == "W" {
			x-=1
		} else if dir == "S" {
			y-=1
		} else {
			fmt.Println( "ERROR, Unhandled direction ["+dir+"]" )
			return
		}
		val = 0
		// Check neighbours (including diagonals) to determine the value for this cell
		cell_val, prs := nodes[Key{ x-1, y-1 }]; if prs { val += cell_val }
		cell_val, prs = nodes[Key{ x, y-1 }]; if prs { val += cell_val }
		cell_val, prs = nodes[Key{ x+1, y-1 }]; if prs { val += cell_val }
		cell_val, prs = nodes[Key{ x-1, y }]; if prs { val += cell_val }
		cell_val, prs = nodes[Key{ x+1, y }]; if prs { val += cell_val }
		cell_val, prs = nodes[Key{ x-1, y+1 }]; if prs { val += cell_val }
		cell_val, prs = nodes[Key{ x, y+1 }]; if prs { val += cell_val }
		cell_val, prs = nodes[Key{ x+1, y+1 }]; if prs { val += cell_val }
		nodes[Key{ x, y }] = val
		// If spot to the 'left' of here is empty then we need to turn left
		// TODO - duplicated following section from above, DRY it somehow
		if dir == "E" { // Check if value to 'left' (N) exists
			_, prs := nodes[Key{ x, y+1 }]
			if !prs { // No value, Turn left (N)
				dir = "N"
			}
		} else if dir == "N" { // Check if value to 'left' (W) exists
			_, prs := nodes[Key{ x-1, y }]
			if !prs { // No value, Turn left (W)
				dir = "W"
			}
		} else if dir == "W" { // Check if value to 'left' (S) exists
			_, prs := nodes[Key{ x, y-1 }]
			if !prs { // No value, Turn left (S)
				dir = "S"
			}
		} else if dir == "S" { // Check if value to 'left' (E) exists
			_, prs := nodes[Key{ x+1, y }]
			if !prs { // No value, Turn left (E)
				dir = "E"
			}
		}
	}
	// Here we should have a value larger than the input number
	fmt.Println( val )
}
