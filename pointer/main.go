package main
import "fmt"

func main() {
	y := [3]int{10, 20, 30}
	fmt.Printf("%v \n", y)
	(*&y)[0] = 100
	fmt.Printf("%v \n", y)
}