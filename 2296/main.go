package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	buildings := make([][3]int, n)
	dp := make([][2]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		input := scanner.Text()
		data := strings.Fields(input)
		x, _ := strconv.Atoi(data[0])
		y, _ := strconv.Atoi(data[1])
		c, _ := strconv.Atoi(data[2])
		buildings[i] = [3]int{x, y, c}
	}

	// x좌표 기준으로 정렬
	sort.Slice(buildings, func(i, j int) bool {
		return buildings[i][0] < buildings[j][0]
	})

	ret := 0
	
  	// 초기 DP 값 설정
	for i := 0; i < n; i++ {
		dp[i][0] = buildings[i][2]
		dp[i][1] = buildings[i][2]
	}

	// DP 테이블 갱신
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
      			// 구간 1,3 경우
			if buildings[j][1] < buildings[i][1] {
				dp[i][0] = max(dp[i][0], dp[j][0]+buildings[i][2])
			}
            
      			// 구간 2,4 경우
			if buildings[j][1] > buildings[i][1] {
				dp[i][1] = max(dp[i][1], dp[j][1]+buildings[i][2])
			}
		}
		// 현재까지의 최대 이익 갱신
		ret = max(ret, max(dp[i][0], dp[i][1]))
	}

	// 결과 출력
	fmt.Fprintln(writer, ret)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
