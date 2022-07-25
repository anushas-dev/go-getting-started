package main

import "fmt"

func main() {
        var x, y string = "foo", "bar"
        x += y
        fmt.Println(x) // foo+bar = foobar

		var p, q int = 27, 7
        p /= q
        fmt.Println(p) // 27/7 = 3

		var a, b float64 = 27.9, 7.0
        a -= b
        fmt.Println(x) // 27.9 - 7.0 = 20.9
        a += b
        fmt.Println(x) // 20.9 + 7.0 = 27.9

		var c, d int = 100,9
        c /= d
        fmt.Println(x) // 100/9 = 11
        c %= d
        fmt.Println(x) // 11%9 = 2

		var e, f int = 100,90
        fmt.Println(e & f)
        fmt.Println(e | f)
        fmt.Println((e+f) >> 2)
}