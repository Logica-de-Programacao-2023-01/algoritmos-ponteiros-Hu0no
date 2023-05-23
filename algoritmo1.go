package main

import "fmt"

func UpdateNumber(ptr *int, n int) int {
        i := 1
        soma := 0
        for i <= n {
                soma += i
                i++
        }
        *ptr = soma
        return *ptr
}
func main() {
        n := 5
        ptr := &n
        fmt.Print(UpdateNumber(ptr, n))
}
