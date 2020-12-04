package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Parse input patches and return all valid passports
func parse(data io.Reader) int {
	// Each passport can span multiple lines
	// k:v pairs separated by spaces on each line
	// Blank line is end of passport
	scanner := bufio.NewScanner(data)
	var (
		hasCid         bool
		validFields    int
		validPassports int
	)
	var (
		validHcl = regexp.MustCompile(`^#[0-9a-f]{6}$`)
		validHgt = regexp.MustCompile(`^(([0-9]{3})(cm)|([0-9]{2})(in))$`)
		validPid = regexp.MustCompile(`^[0-9]{9}$`)
	)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			if validFields == 7 && !hasCid || validFields == 8 {
				validPassports++
			}
			// Move to next entry; reset values
			hasCid = false
			validFields = 0
			continue
		}

		for _, part := range strings.Split(t, " ") {
			field := strings.Split(part, ":")
			switch field[0] {
			case "byr":
				yr, err := strconv.Atoi(field[1])
				if err != nil {
					// fmt.Fprintln(os.Stderr, "invalid birth year value", field[1])
					continue
				}
				if yr >= 1920 && yr <= 2002 {
					validFields++
				}
			case "iyr":
				yr, err := strconv.Atoi(field[1])
				if err != nil {
					// fmt.Fprintln(os.Stderr, "invalid issue year value", field[1])
					continue
				}
				if yr >= 2010 && yr <= 2020 {
					validFields++
				}
			case "eyr":
				yr, err := strconv.Atoi(field[1])
				if err != nil {
					// fmt.Fprintln(os.Stderr, "invalid expiration year value", field[1])
					continue
				}
				if yr >= 2020 && yr <= 2030 {
					validFields++
				}
			case "hgt":
				m := validHgt.FindStringSubmatch(field[1])
				if len(m) == 0 {
					continue
				}
				switch {
				case m[3] == "cm":
					h, _ := strconv.Atoi(m[2])
					if h >= 150 && h <= 193 {
						validFields++
					}
				case m[5] == "in":
					h, _ := strconv.Atoi(m[4])
					if h >= 59 && h <= 76 {
						validFields++
					}
				}
			case "hcl":
				if validHcl.MatchString(field[1]) {
					validFields++
				}
			case "ecl":
				switch field[1] {
				case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
					validFields++
				default:
					continue
				}
			case "pid":
				if validPid.MatchString(field[1]) {
					validFields++
				}
			case "cid":
				validFields++
				hasCid = true
			default:
				fmt.Fprintln(os.Stderr, "unknown field", field[0])
			}
		}
	}
	// Recheck after EOF for last entry
	if validFields == 7 && !hasCid || validFields == 8 {
		validPassports++
	}
	return validPassports
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open input: %v", err)
	}

	p := parse(f)
	fmt.Println("Valid:", p)
}
