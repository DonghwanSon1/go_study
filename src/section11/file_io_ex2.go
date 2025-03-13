package main

import (
	"bufio"
	"fmt"
	"os"
)

func errCheck9(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 파일 읽기, 버퍼사용 쓰기 -> bufio 패키지 활용
	// ioutil, bufio 등은 Io.Reader, io.Writer 인터페이스를 구현 함
	//	-> 즉, Writer, Read 메서드를 사용 가능
	//  -> 즉, bufio 의 NewReader, NewWriter 를 통해서 객체 반환 후 메서드 사용

	// bufio(Buffered io) 패키지
	// https://golang.org/pkg/bufio

	// 파일 열기
	// 두번째 매개변수 확인
	// https://golang.org/pkg/os/#pkg-constants

	// 파일 생성
	file, err := os.OpenFile("test_write2.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	errCheck9(err)

	// bufio 파일 쓰기 예제
	wt := bufio.NewWriter(file)
	wt.WriteString("Hello golang! \nFile Write Test!!! - 1\n")
	wt.Write([]byte("Hello golang! \nFile Write Test!!! - 2\n")) // 여기까지는 버퍼에 가지고 있다. 즉, 파일에는 아무것도 안써져있고, 고랭 내부적인 버퍼에 가지고 있는거다.

	fmt.Printf("사용한 버퍼 사이즈 : %d Byte\n", wt.Buffered())
	fmt.Printf("남은 버퍼 사이즈 : %d Byte\n", wt.Available())
	fmt.Printf("전체 버퍼 사이즈 : %d Byte\n", wt.Size())

	wt.Flush() // 해당 함수를 통해 버퍼를 초기화 하고 파일 쓰기를 진행하게 된다. 즉, 버퍼는 비워지고 파일안에 값이 넣어진다.

	fmt.Println("------------- 쓰기 작업 완료 -------------")

	rt := bufio.NewReader(file)
	fi, err := file.Stat()
	errCheck9(err)

	b := make([]byte, fi.Size())
	fmt.Println("파일 정보 출력 : ", fi)
	fmt.Println("파일 정보 이름 : ", fi.Name())
	fmt.Println("파일 정보 크기 : ", fi.Size())
	fmt.Println("파일 정보 수정시간 : ", fi.ModTime())
	fmt.Println("------------------------------------")

	file.Seek(0, os.SEEK_SET)
	data, err := rt.Read(b)
	//data, err := ioutil.ReadFile(fi.Name()) // 다른 방법 (file.Seek 와 b 변수 없이)
	errCheck9(err)

	fmt.Printf("전체 Buffer Size : %d Bytes \n", rt.Size())
	fmt.Printf("읽기 작업 완료 : %d Bytes \n", data)
	fmt.Println("------------------------------------")
	fmt.Println(string(b))
	fmt.Println("------------------------------------")

	defer file.Close()

}
