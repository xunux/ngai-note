package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	log.Print("Hello, ", "世界")

	var integer32 int = 2147483648
	fmt.Println(integer32)
	fmt.Println(math.MaxInt)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MaxInt == math.MaxInt64)
	var b bool = true
	fmt.Println(b)
	var s string = "xiao"
	fmt.Println(s)
	var result, result2 = calc("4", "3")
	fmt.Println(result, result2)
	if result == result2 {
		fmt.Println(true)
	}

	switch {
	case integer32 > 1000:
		fmt.Println("integer32 > 1000")
		break
	case integer32 > 100:
		fmt.Println("integer32 > 100")
		break
	default:
		fmt.Println("integer32", integer32)
	}

	var arr [3]int
	fmt.Println(arr)

	var strs = [...]string{"ni", "wo", "ta"}
	fmt.Println(strs)

	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("main(): recover", handler)
		}
	}()

	highlow(2, 0)
	fmt.Println("Program finished successfully!")

}

func calc(number1 string, number2 string) (sum int, mul int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	sum = int1 + int2
	mul = int1 * int2
	fmt.Println(number1, number2, int1, int2, sum, mul)
	return sum, mul
}

func highlow(high int, low int) {
	if high < low {
		fmt.Println("Panic!")
		panic("highlow() low greater than high")
	}
	defer fmt.Println("Deferred: highlow(", high, ",", low, ")")
	fmt.Println("Call: highlow(", high, ",", low, ")")

	highlow(high, low+1)
}
