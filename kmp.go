package main

import (
    "fmt"
)

func strKMP(t, p string) []int {
    lPrefix := func(p string) []int {
        lp := make([]int, len(p))
        for k, i := 0, 1; i < len(p); i++ {
            for k > 0 && p[k] != p[i] {
                k = lp[k - 1]
            }
            if p[k] == p[i] {
                k++
            }
            lp[i] = k
        }
        return lp
    }(p)

    match, count := make([]int, len(t)), 0
    for x, y := 0, 0; x < len(t); x++ {
        for y > 0 && p[y] != t[x] {
            y = lPrefix[y - 1]
        }
        if p[y] == t[x] {
            y++
        }
        if y == len(p) {
            match[count] = x - len(p) + 1
            count += 1
            y = lPrefix[y - 1]
        }
    }

    return match[:count]
}

func main() {
    var test, pattern string
    fmt.Println("Enter Test String: ")
    fmt.Scanln(&test)
    fmt.Println("Enter Pattern String: ")
    fmt.Scanln(&pattern)
    match := strKMP(test, pattern)
    fmt.Printf("%d match found: \n", len(match))
    for i := range match {
        fmt.Printf("[%d]: %d\n", i + 1, match[i])
    }
}
