package main

import (
	"fmt"
	"time"
)

func exe(name string) {
	fmt.Println(name, " Start !! -> ", time.Now())
	for i := 0; i < 1000; i++ {
		fmt.Println(name, " >>>>>> ", i)
	}
	fmt.Println(name, " End !! -> ", time.Now())
}

func main() {
	// 고루틴
	exe("t1")

	// 예제 1
	fmt.Println("Main Routine Start!! -> ", time.Now())
	go exe("t2")
	go exe("t3")
	time.Sleep(4 * time.Second)

}
