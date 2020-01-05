package main
 
import (
	"fmt"
)
 
func main() {
	var (
		n     int
		count int
	)
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
 
	flag := true
	for flag {
		for i := 0; i < n; i++ {
			if nums[i]%2 == 0 {
				nums[i] /= 2
			} else {
				flag = false
			}
		}
		if flag {
			count++
		}
	}
	fmt.Println(count)
}
