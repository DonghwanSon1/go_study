package main

import "fmt"

func main() {
	// 구조체 익명 선언 및 활용
	// 예제 1
	car1 := struct {
		name, color string
	}{"520d", "red"}

	fmt.Println("ex 1 : ", car1)
	fmt.Printf("ex 1 : %#v\n", car1)

	// 예제 2
	cars := []struct{ name, color string }{
		{"520d", "red"},
		{"530i", "white"},
		{"528i", "blue"},
	}

	for _, car := range cars {
		fmt.Printf("(%s, %s) ----- (%#v)\n", car.name, car.color, car)
	}

}
