package chapter3

import "fmt"

func Chap3main() {
	fmt.Println("1+1=", 1+1)

	// fmt.Println(fmt.Sprintf("%.1f", 1.0+1.0))
	fmt.Println("1.0 + 1.0 =", fmt.Sprintf("%.1f", 1.0+1.0))
}
