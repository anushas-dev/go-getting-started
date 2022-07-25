package main
import "fmt"

const PI float64 = 3.14 // global constant

func main() {
    var radius float64 = 5.154
    var area float64

    area = PI * radius * radius
    fmt.Printf("Radius: %.2f \nPI:%.1f \n", radius, PI)
    fmt.Printf("Area of Circle is : %f \n", area)
    fmt.Println("Thank You")
}