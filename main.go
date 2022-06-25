package main

import "fmt"

type interval struct {
	start int
	end   int
}

func main() {
	// 3, 5
	slice := []int{-3, 5, -20, 10, 1, 5, -30}
	l, h := GetHighestSumInterval(slice)
	fmt.Printf("\n\nthe highest sum is between %d and %d", l, h)
}

func GetHighestSumInterval(slice []int) (int, int) {
	highestSum := slice[0]
	var resp interval
	for i := 0; i < len(slice); i++ {
		sum := slice[i]
		fmt.Printf("%d: ", i)
		fmt.Printf("highest sum: %d\n", highestSum)
		for j := i + 1; j < len(slice); j++ {
			newSum := sum + slice[j]
			fmt.Printf("last sum: %d\n", sum)
			fmt.Printf("new sum: %d\n", newSum)
			if newSum < sum {
				break
			}
			sum = newSum
			if sum > highestSum {
				highestSum = sum
				resp.start = i
				resp.end = j
			}
		}
	}
	return resp.start, resp.end
}
