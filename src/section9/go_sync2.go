package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 구조체 선언 (공유 데이터)
type count1 struct {
	num   int
	mutex sync.Mutex
}

func main() {
	// 고루틴 동기화 예제
	// 실행 흐름 제어 및 변수 동기화 가능
	// 공유 데이터 보호가 가장 중요

	// 동기화 사용 예제 - Mutex : 여러 고루틴에서 작업하는 공유데이터 보호
	// sync.Mutex 선언 후 Lock, 끝나면 Unlock 사용 한다.
	// 시스템 전체 cpu 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count1{num: 0}
	done := make(chan bool)

	for i := 1; i <= 10000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched() // CPU 양보
		}()
	}

	for i := 1; i <= 10000; i++ {
		<-done
	}

	c.result1()
}

func (c *count1) increment() {
	// 공유 데이터 수정 전 뮤텍스로 보호
	c.mutex.Lock()
	c.num += 1
	// 공유 데이터 수정 후 보호 해제
	c.mutex.Unlock()
}

func (c *count1) result1() {
	fmt.Println("result : ", c.num)
}
