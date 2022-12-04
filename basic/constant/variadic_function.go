package main

func sum(nums ...int) int { // nums จะเป็น slice
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	println(sum(1, 2))

	nums := []int{1, 2, 3, 4}
	println(sum(nums...))
}
