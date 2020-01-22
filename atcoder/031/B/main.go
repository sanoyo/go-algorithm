// https://atcoder.jp/contests/arc031/submissions/5810403
package main

import "fmt"

var (
	C      [10][10]rune
	N      int // the number of lands
	marked int
	dy     = []int{1, 0, -1, 0}
	dx     = []int{0, 1, 0, -1}
)

func main() {
	for i := 0; i < 10; i++ {
		var s string
		fmt.Scan(&s)
		for j := 0; j < 10; j++ {
			C[i][j] = rune(s[j])
			if C[i][j] == 'o' {
				N++
			}
		}
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if C[i][j] == 'o' {
				continue
			}
			C[i][j] = 'o'
			if dfs(i, j) {
				fmt.Println("YES")
				return
			}
			C[i][j] = 'x'
			reset()
		}
	}
	fmt.Println("NO")
}

func dfs(y, x int) bool {
	if marked == N {
		return true
	}
	C[y][x] = '#'
	marked++
	for i := 0; i < len(dy); i++ {
		ny := y + dy[i]
		nx := x + dx[i]
		if 0 <= nx && nx < 10 && 0 <= ny && ny < 10 && C[ny][nx] == 'o' {
			if dfs(ny, nx) {
				return true
			}
		}
	}
	return false
}

func reset() {
	marked = 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if C[i][j] == '#' {
				C[i][j] = 'o'
			}
		}
	}
}
