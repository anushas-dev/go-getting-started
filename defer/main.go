package main

import "fmt"

func printString(str string){
        fmt.Printf("%q ", str)
}

func printInt(i int){
        fmt.Printf("%d ", i)
}

func printFloat(f float64){
        fmt.Printf("%.2f ", f)
}
func main() {
        printString("browser")
        defer printInt(32)
        defer printFloat(0.24)
        printString("chrome")
        printInt(90)
        defer printFloat(89)
        printInt(900)
}