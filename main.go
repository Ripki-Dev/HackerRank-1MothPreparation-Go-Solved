package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println(Solution(24, 12))
	numb := []int32{1, 2, 3, 4, 5}
	miniMaxSum(numb)
}

func Solution(n int, m int) int {
	// Your code starts here.
	var res []int
	var high, low int

	if n <= m {
		high = m
		low = n
	} else {
		high = n
		low = m
	}
	go func() {
		for i := 1; i <= high; i++ {
			if n == m {
				if n%i == 0 && m%i == 0 {
					res = append(res, i)
				}
			} else if high%low == 0 {
				if n%i == 0 && m%i == 0 && i != high {
					res = append(res, i)
				}
			} else {
				if n%i == 0 && m%i == 0 && i != n && i != m {
					res = append(res, i)
				}
			}

		}

	}()
	time.Sleep(500 * time.Millisecond)
	return len(res)
}

func plusMinus(arr []int32) {
	// Write your code here
	var positive, negative, zero int
	devider := len(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			positive++
		} else if arr[i] < 0 {
			negative++
		} else {
			zero++
		}
	}
	fmt.Println(float64(positive) / float64(devider))
	fmt.Println(float64(negative) / float64(devider))
	fmt.Println(float64(zero) / float64(devider))

}

func miniMaxSum(arr []int32) {
	// Write your code here
	var res []int64
	for i := 0; i < len(arr); i++ {
		var nitip int64
		for j := 0; j < len(arr); j++ {
			if i != j {
				nitip += int64(arr[j])
			}
		}
		res = append(res, nitip)
	}
	min := res[0]
	max := res[0]
	for _, value := range res {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	fmt.Println(max)
	fmt.Println(min)
}
