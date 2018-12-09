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

type fabric map[coord][]int

func coverage(claims []claim) (int, int) {
	var inches int
	fab := make(fabric)
	uniq := make(map[int]struct{})
	for _, c := range claims {

		for x := c.Left; x <= c.Right; x++ {
			for y := c.Top; y <= c.Bottom; y++ {
				fab[coord{x, y}] = append(fab[coord{x, y}], c.ID)
			}

		}

		uniq[c.ID] = struct{}{}

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

	for _, c := range fab {
		if len(c) > 1 {
			inches++
		}
	}

	for _, ids := range fab {
		if len(ids) <= 1 {
			continue
		}
		for _, id := range ids {
			delete(uniq, id)
		}
		if len(uniq) == 1 {
			break
		}
	}
	var uc int
	for id := range uniq {
		uc = id
		break // There should only be one entry
	}

	return inches, uc
}

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	sqin, noOverlap := coverage(allClaims(input))
	fmt.Printf("Part one: %d sq. in.\n", sqin)
	fmt.Printf("Part two: ID %d\n", noOverlap)
}
