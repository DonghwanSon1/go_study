package main

// 선언 방식 1
//import "fmt"
//import "os"

// 선언 방식 2
import (
	"fmt"
	"os"
)

func main() {
	// 패키지 : 코드 모듈화 및 재사용
	// 응집도 높이고, 결합도 낮추는 재사용성이 가능한 클린 코드
	// golang : 패키지 단위의 독립적이고 작은 단위로 개발 -> 작은 패키지를 결합해서 프로그램을 작성할 것을 권고 한다.
	// 패키지 이름 = 디렉토리 이름
	// 같은 패키지 내의 소스파일들은 디렉토리명을 패키지 명으로 사용한다.
	// 네이밍 규칙 : 소문자는 private, 대문자는 public 이다.
	// golang : main 패키지는 특별하게 인식 된다. -> 따라서 main 함수를 가지고 있다면, 패키지명은 Main 이여야 하며 컴파일러 공유 라이브러리 X, 프로그램의 시작점 start point 이다.

	var name string

	fmt.Println("이름은 ? : ")
	fmt.Scanf("%s", &name)

	fmt.Fprintf(os.Stdout, "Hi! %s\n", name)

}
