package main

import "fmt"

func main() {
	// 여러개 선언
	var (
		name      string = "machine"
		height    int32
		weight    int32
		isRunning bool
	)

	height = 250
	weight = 350
	isRunning = true

	fmt.Println("name : ", name, ", height : ", height, ", weight : ", weight, ", isRunning : ", isRunning)
}
