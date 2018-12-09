package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type claim struct {
	ID     int
	Left   int
	Top    int
	Bottom int
	Right  int
	Width  int
	Height int
}

func str2int(s []string) []int {
	var ints []int
	for _, v := range s[1:] {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}

func newClaim(id string) claim {
	re := regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)
	match := str2int(re.FindStringSubmatch(id))
	return claim{
		ID:     match[0],
		Left:   match[1],
		Top:    match[2],
		Bottom: match[2] + match[4] - 1,
		Right:  match[1] + match[3] - 1,
		Width:  match[3],
		Height: match[4],
	}
}

func allClaims(input []string) []claim {
	var claims []claim
	for _, line := range input {
		claims = append(claims, newClaim(line))
	}
	return claims
}

type coord struct {
	x, y int
}

type fabric map[coord]int

func coverage(claims []claim) int {
	var inches int
	claimed := make(fabric)
	for _, c := range claims {
		for x := c.Left; x <= c.Right; x++ {
			for y := c.Top; y <= c.Bottom; y++ {
				claimed[coord{x, y}]++
			}
		}

		// Just look at all the complicated things I tried before the above 😭

		// for j, b := range claims[i+1:] {
		// 	if j == len(claims)-1 {
		// 		break
		// 	}
		// var overlap int

		// Wow I'm doing this the hard way.
		// if a.Left <= b.Right && b.Left <= a.Left {
		// 	// A is horizontally inside B
		// 	if a.Top <= b.Bottom && b.Bottom <= a.Bottom {
		// 		// A is vertically inside B
		// 		overlap := (b.Width - a.Width) * (b.Height - a.Height)
		// 		inches += overlap
		// 		fmt.Printf("A: %#v\nB: %#v\n", a, b)
		// 		fmt.Printf("Inside A:%d B:%d W:%d H:%d C:%d\n", a.ID, b.ID, b.Width-a.Width, b.Height-a.Height, overlap)
		// 	}
		// } else if a.Left <= b.Left && b.Left <= a.Right {
		// 	// A overlaps B horizontally on the left
		// 	if a.Top <= b.Bottom && b.Bottom <= a.Bottom {
		// 		// A overlaps B horizontally
		// 		overlap := (a.Right - (b.Left - 1)) * (b.Bottom - (a.Top - 1))
		// 		inches += overlap
		// 		fmt.Printf("A: %#v\nB: %#v\n", a, b)
		// 		fmt.Printf("Overlap A:%d B:%d W:%d H:%d C:%d\n", a.ID, b.ID, a.Right-(b.Left-1), b.Bottom-(a.Top-1), overlap)
		// 	}
		// }

		// if a.Right >= b.Left && b.Bottom >= a.Top && a.Bottom >= b.Top {
		// 	overlap := (a.Right - (b.Left - 1)) * (b.Bottom - (a.Top - 1))
		// 	if overlap > inches {
		// 		inches = overlap
		// 	}
		// 	fmt.Printf("A: %#v\nB: %#v\n", a, b)
		// 	fmt.Printf("A:%d B:%d W:%d H:%d C:%d\n", a.ID, b.ID, a.Right-(b.Left-1), b.Bottom-(a.Top-1), overlap)
		// }
		// }
	}

	for _, c := range claimed {
		if c > 1 {
			inches++
		}
	}
	return inches
}

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Printf("Part one: %d sq. in.\n", coverage(allClaims(input)))
}
