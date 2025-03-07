package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 예제 1
	rand.Seed(time.Now().UnixNano())
	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println(i, " -> i값, ", "50 이상이면서 100 미만입니다.")
	case i >= 25 && i < 50:
		fmt.Println(i, " -> i값, ", "25 이상이면서 50 미만입니다.")
	default:
		fmt.Println(i, " -> i값, ", "25 미만입니다.")
	}

}
