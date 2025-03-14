package main

/**
 * fmt - formatting & printing
 * math/rand - generating random keys
 * net/http - handling http requests
 * time - seeding the random number generator
 */

// Hello world
/* func main() {
 	fmt.Println("hello world")
} */

// Values
/* func main() {
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
} */

// Variables
/* func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
} */

// Constants
/* import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
} */

// For
/* import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
} */

// If/Else
/* import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
} */
