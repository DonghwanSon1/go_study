package main

import (
	"fmt"
	"section4/lib"
)

func main() {
	// 현재 Goland에서 실행 시 go mod 로 진행하기 때문에 해당 패키지를 인식 못함
	// 하지만 강의에서는 GOPATH를 통해서 해당 내용을 진행하고 있기 때문에
	// goland 설정에 들어가서 go 모듈에서 'go 모듈 통합 활성화' 를 꺼준다.
	fmt.Println("10 보다 큰 수 ? : ", lib.CheckNum(11))

}
