package main

import (
	"fmt"
	"reflect"
)

type Car1 struct {
	name    string "차량명"
	color   string "색상"
	company string "제조사"
}

func main() {
	// 필드 태그 사용
	// 예제 1
	tag := reflect.TypeOf(Car1{})

	for i := 0; i < tag.NumField(); i++ {
		fmt.Println("ex 1 : ", tag.Field(i).Tag, tag.Field(i).Name, tag.Field(i).Type)
	}

}
