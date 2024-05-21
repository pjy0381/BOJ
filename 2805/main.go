package main

import (
    "bufio"
    "fmt"
    "os"
)

var (
    arr    []int
    n, m   int
    answer int
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Fscanf(reader, "%d %d\n", &n, &m)
    arr = make([]int, n)
    
    var min, max int
    for i := 0; i < n; i++ {
        fmt.Fscanf(reader, "%d", &arr[i])
        if arr[i] > max {
            max = arr[i]
        }
    }

    solution(min, max)
    fmt.Println(answer)
}

func solution(min, max int) {
    if min > max {
        answer = max
        return
    }
    var result int64
    mid := (min + max) / 2
    for i := 0; i < len(arr); i++ {
        if arr[i] > mid {
            result += int64(arr[i] - mid)
        }
    }
    if result >= int64(m) {
        solution(mid+1, max)
    }
    if result < int64(m) {
        solution(min, mid-1)
    }
}
