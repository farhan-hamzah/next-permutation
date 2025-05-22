package main

import (
	"fmt"
)

func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	// Step 1: Cari indeks 'i' dari belakang, di mana nums[i] < nums[i+1]
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		// Step 2: Cari indeks 'j' dari belakang, di mana nums[j] > nums[i]
		j := n - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		// Step 3: Tukar nums[i] dan nums[j]
		nums[i], nums[j] = nums[j], nums[i]
	}

	// Step 4: Balik urutan dari nums[i+1] sampai akhir
	reverse(nums, i+1, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// Contoh penggunaan
func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums) // Output: [1 3 2]
}
