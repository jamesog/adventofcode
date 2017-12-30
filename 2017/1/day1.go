package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func sumNext(nums []int) int {
	var total int

	for i := 0; i < len(nums); i++ {
		if i < len(nums)-1 {
			if nums[i] == nums[i+1] {
				total += nums[i]
			}
		} else {
			if nums[i] == nums[0] {
				total += nums[i]
			}
		}
	}

	return total
}

func sumHalfway(nums []int) int {
	var total int

	halfway := len(nums) / 2

	// fmt.Println(nums)
	// fmt.Println("len:", len(nums))
	for i, _ := range nums {
		n := (i + halfway)
		if i+halfway >= len(nums) {
			n -= len(nums)
		}
		// fmt.Printf("i: %d=%d, halfway: %d=%d\n", i, nums[i], n, nums[n])
		if nums[i] == nums[n] {
			total += nums[i]
		}
	}
	return total
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", input)
	// fmt.Println(string(input[0]))
	var nums []int
	// var f, p int

	for _, b := range input {
		if string(b) == "\n" {
			continue
		}
		c, err := strconv.Atoi(string(b))
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, c)
	}

	total := sumNext(nums)
	fmt.Println(total)
	total = sumHalfway(nums)
	fmt.Println(total)
}
