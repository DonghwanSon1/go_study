package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 고루틴 동기화 고급
	// Once : 한번만 실행(주로 초기화에 사용)
	// Do로 실행

	runtime.GOMAXPROCS(runtime.NumCPU())

	once := new(sync.Once)
	for i := 0; i < 5; i++ {
		go func(n int) {
			once.Do(onceTest)
			fmt.Println("Goroutine : ", n)
		}(i)
	}

	time.Sleep(2 * time.Second)
}

func onceTest() {
	// 이 부분에 한번 실행 할 코드 작성
	fmt.Println("Once Test Execute!!")
}
