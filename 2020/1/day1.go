package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sum2(nums []int) (int, int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == 2020 {
				return nums[i], nums[j]
			}
		}
	}
	return -1, -1
}

func sum3(nums []int) (int, int, int) {
	for i := 0; i < len(nums)-2; i++ {
		for j := 1; j < len(nums)-1; j++ {
			for k := 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 2020 {
					return nums[i], nums[j], nums[k]
				}
			}
		}
	}
	return -1, -1, -1
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open input: %v", err)
	}
	var nums []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}

	num1, num2 := sum2(nums)
	part1 := num1 * num2
	fmt.Println("Part 1:", part1)

	num1, num2, num3 := sum3(nums)
	part2 := num1 * num2 * num3
	fmt.Println("Part 2:", part2)
}
